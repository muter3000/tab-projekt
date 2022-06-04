package auth

import (
	"github.com/tab-projekt-backend/database/redis"
)

// Registered error coming from bad client request
// swagger:response error400
type responseError400 struct {
	//in: body
	Body string
}

// Registered error coming from internal server error
// swagger:response error500
type responseError500 struct {
	//in: body
	Body string
}

// Session creation success
// swagger:response success
type responseSuccess struct {
	//in: cookie
	//name: session-id
	SessionID string
}

// swagger:parameters CreateSession
type levelWrapper struct {
	// The wanted access level: 1-Pracownik 2-Administrator 3-AdministratorDB
	// in: path
	// required: true
	Id redis.PermissionLevel `json:"level"`
}

// swagger:route POST /auth/{level} Auth CreateSession
// Returns list of administratorzy
// responses:
//  201: success
//  400: error400
//  500: error500
