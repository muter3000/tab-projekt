// Package classification Workers API
//
// Documentation of Workers API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  -application/json
//
//  Produces:
//  -application/json
// swagger:meta
package handlers

import "github.com/gorilla/mux"

type SubRouter interface {
	// RegisterSubRouter Funkcja rejestruje subrouter dla podanego routera
	RegisterSubRouter(router *mux.Router)
}
