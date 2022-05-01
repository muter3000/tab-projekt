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

func NewRedisClient(l hclog.Logger) *RedisClient {
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
		l.Error("pinging regis", "err", err)
		return nil
	}

	return &RedisClient{l: l, client: rdb}
}

func (rc *RedisClient) CreateUserSession(rw http.ResponseWriter, r *http.Request, level PermissionLevel) {
	sId := uuid.New().String()
	cookie := &http.Cookie{
		Name:   "session-id",
		Value:  sId,
		MaxAge: 300,
	}
	r.AddCookie(cookie)

	err := rc.client.Set(context.Background(), sId, level, 300*time.Second).Err()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rc *RedisClient) CheckAuthorization(r *http.Request, level PermissionLevel) bool {
	response := &http.Response{
		Request: r,
	}
	cookies := response.Cookies()

	sId := ""
	for _, c := range cookies {
		if c.Name == "session-id" {
			sId = c.Value
			break
		}
	}
	if sId == "" {
		return false
	}
	value, err := rc.client.Get(context.Background(), sId).Result()
	if err != nil {
		rc.l.Error("getting session from redis", "err", err)
		return false
	}
	val, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	if PermissionLevel(val) != level {
		return false
	}
	return true
}
