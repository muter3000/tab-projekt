package redis

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/schemas"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type PermissionLevel int8

const (
	None PermissionLevel = iota
	Kierowca
	Administator
	AdministratorDB
)

type AuthorizationClient struct {
	l  hclog.Logger
	rc *redis.Client
	pg *pg.DB
}

type Session struct {
	SessionID string
	Level     PermissionLevel
}

func NewAuthorizationClient(l hclog.Logger, pg *pg.DB) (*AuthorizationClient, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	rdb := redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", host, port),
		Password:           "", // no password set
		DB:                 0,  // use default DB
		IdleTimeout:        time.Second * 25,
		IdleCheckFrequency: time.Second * 5,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &AuthorizationClient{l: l, rc: rdb, pg: pg}, nil
}

func (a *AuthorizationClient) CreateUserSession(rw http.ResponseWriter, login string, password string, level PermissionLevel) bool {
	a.l.Debug("Handling session creation with", "login", login, "level", level)
	switch level {
	case Kierowca:
		kierowca := schemas.Kierowca{}
		err := a.pg.Model(&kierowca).Where("login = ?", login).Select()
		if err != nil {
			a.l.Debug("err happened while reading from db", "err", err)
			return false
		}
		if err = bcrypt.CompareHashAndPassword([]byte(kierowca.Haslo), []byte(password)); err != nil {
			a.l.Debug("bcrypt.Compare", "hash", kierowca.Haslo, "pass", password, "err", err)
			return false
		}
	case Administator:
		var administrator schemas.Administrator
		err := a.pg.Model(&administrator).Where("login = ?", login).Select()
		if err != nil {
			a.l.Debug("err happened while reading from db", "err", err)
			return false
		}
		if err = bcrypt.CompareHashAndPassword([]byte(administrator.Haslo), []byte(password)); err != nil {
			a.l.Debug("bcrypt.Compare", "hash", administrator.Haslo, "pass", password, "err", err)
			return false
		}
	case AdministratorDB:
		if login != "adminDB" || password != os.Getenv("DB_ADMIN_PASS") {
			return false
		}
	}

	sId := uuid.New().String()
	cookie := &http.Cookie{
		Name:   "session-id",
		Value:  sId,
		MaxAge: 300,
	}
	http.SetCookie(rw, cookie)

	err := a.rc.Set(context.Background(), sId, strconv.Itoa(int(level)), 300*time.Second).Err()
	if err != nil {
		return false
	}
	a.l.Debug("Created session with", "id", sId, "level", int8(level))
	return true
}

func (a *AuthorizationClient) InvalidateUserSession(r *http.Request) bool {
	cookies := r.Cookies()
	sId := ""
	for _, c := range cookies {
		if c.Name == "session-id" {
			sId = c.Value
			break
		}
	}
	if sId == "" {
		a.l.Debug("user didn't send session cookie")
		return false
	}
	err := a.rc.Del(context.Background(), sId).Err()
	if err != nil {
		a.l.Warn("db error while invalidating token", "err", err)
		return false
	}
	a.l.Debug("Invalidated session", "sessionID", sId)
	return true
}

func (a *AuthorizationClient) CheckAuthorization(r *http.Request, level PermissionLevel) bool {
	cookies := r.Cookies()

	sId := ""
	for _, c := range cookies {
		if c.Name == "session-id" {
			sId = c.Value
			break
		}
	}
	if sId == "" {
		a.l.Debug("user didn't send session cookie")
		return false
	}
	value, err := a.rc.Get(context.Background(), sId).Result()
	if err != nil {
		a.l.Warn("getting session from redis", "err", err)
		return false
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		a.l.Error("converting val to int", "val", val)
		return false
	}
	if PermissionLevel(val) != level {
		a.l.Debug("Got session with wrong permission level", "uuid", sId, "permissionLevel", val, "requiredLevel", int8(level))
		return false
	}
	a.l.Debug("Got session", "uuid", sId, "permissionLevel", val)
	return true
}
