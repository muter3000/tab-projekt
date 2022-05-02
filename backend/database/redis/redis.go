package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
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
	AdminstratorDB
)

type RedisClient struct {
	l      hclog.Logger
	client *redis.Client
}

type Session struct {
	SessionID string
	Level     PermissionLevel
}

func NewRedisClient(l hclog.Logger) (*RedisClient, error) {
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

	return &RedisClient{l: l, client: rdb}, nil
}

func (rc *RedisClient) CreateUserSession(rw http.ResponseWriter, level PermissionLevel) bool {
	sId := uuid.New().String()
	cookie := &http.Cookie{
		Name:   "session-id",
		Value:  sId,
		MaxAge: 300,
	}
	http.SetCookie(rw, cookie)

	err := rc.client.Set(context.Background(), sId, strconv.Itoa(int(level)), 300*time.Second).Err()
	if err != nil {
		return false
	}
	rc.l.Debug("Created session with", "id", sId, "level", int8(level))
	return true
}

func (rc *RedisClient) InvalidateUserSession(rw http.ResponseWriter, r *http.Request) bool {
	cookies := r.Cookies()
	sId := ""
	for _, c := range cookies {
		if c.Name == "session-id" {
			sId = c.Value
			break
		}
	}
	if sId == "" {
		rc.l.Debug("user didn't send session cookie")
		return false
	}
	err := rc.client.Del(context.Background(), sId).Err()
	if err != nil {
		rc.l.Warn("db error while invalidating token", "err", err)
		return false
	}
	rc.l.Debug("Invalidated session", "sessionID", sId)
	return true
}

func (rc *RedisClient) CheckAuthorization(r *http.Request, level PermissionLevel) bool {
	cookies := r.Cookies()

	sId := ""
	for _, c := range cookies {
		if c.Name == "session-id" {
			sId = c.Value
			break
		}
	}
	if sId == "" {
		rc.l.Debug("user didn't send session cookie")
		return false
	}
	value, err := rc.client.Get(context.Background(), sId).Result()
	if err != nil {
		rc.l.Warn("getting session from redis", "err", err)
		return false
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		rc.l.Error("converting val to int", "val", val)
		return false
	}
	if PermissionLevel(val) != level {
		rc.l.Debug("Got session with wrong permission level", "uuid", sId, "permissionLevel", val, "requiredLevel", int8(level))
		return false
	}
	rc.l.Debug("Got session", "uuid", sId, "permissionLevel", val)
	return true
}
