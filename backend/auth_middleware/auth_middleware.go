package auth_middleware

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/redis"
	"net/http"
	"os"
)

type AuthorizationMiddleware struct {
	l hclog.Logger
	Authorizer
}

func NewAuthorisationMiddleware(l hclog.Logger, authorizer Authorizer) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{l: l, Authorizer: authorizer}
}

func (ah *AuthorizationMiddleware) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookies := request.Cookies()
		sId := ""
		for _, c := range cookies {
			if c.Name == "session-id" {
				sId = c.Value
				break
			}
		}
		if sId == "" {
			http.Error(writer, "The request could not be authorized", http.StatusUnauthorized)
			ah.l.Info("Request didn't send session id")
			return
		}

		_, err := ah.CheckAuthorization(request)
		if err != nil {
			http.Error(writer, "The request could not be authorized", http.StatusUnauthorized)
			ah.l.Info("Request couldn't be authorized", "err", err)
			return
		}

		if err != nil {
			http.Error(writer, "The request could not be authorized", http.StatusUnauthorized)
			ah.l.Info("Request couldn't be authorized", "err", err)
			return
		}
		h.ServeHTTP(writer, request)
	})
}

type Authorizer struct {
	Level redis.PermissionLevel
}

func (k Authorizer) CheckAuthorization(r *http.Request) (bool, error) {
	authorization := getAuthorization(r, k.Level)
	if authorization != nil {
		return false, authorization
	}
	return true, nil
}

func getAuthorization(r *http.Request, level redis.PermissionLevel) error {
	authHost := os.Getenv("AUTH_HOST")
	authPort := os.Getenv("AUTH_PORT")

	url := fmt.Sprintf("%s:%s/auth/%v", authHost, authPort, int8(level))

	req, err := http.NewRequest(http.MethodGet, url, nil)

	for _, c := range r.Cookies() {
		req.AddCookie(c)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request for %v level unauthorized", level)
	}
	return nil
}
