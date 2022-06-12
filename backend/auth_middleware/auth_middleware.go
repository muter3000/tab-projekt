package auth_middleware

import (
	_ "github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/tab-projekt-backend/database/redis"
	"github.com/tab-projekt-backend/helpers"
	"net/http"
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
	_, authorization := helpers.GetAuthAndIDFromSession(r, k.Level)
	if authorization != nil {
		return false, authorization
	}
	return true, nil
}
