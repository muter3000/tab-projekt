package administratorzy

import "github.com/tab-projekt-backend/schemas"

// Registered error
// swagger:response error
type responseError struct {
	// in: body
	Body string
}

// swagger:route GET /administratorzy/ Administrators ListAdministrators
// Returns list of administrators
// responses:
//  200: administratorsResponse
//  500: error

// A list of all registered administrators
// swagger:response administratorsResponse
type administratorsResponseWrapper struct {
	// All registered administrators
	// in: body
	Body []schemas.Administrator
}

// swagger:parameters GetAdministratorByID
type idWrapper struct {
	// The id of the wanted administrator
	// in: path
	// required: true
	Id int32 `json:"id"`
}

// swagger:route GET /administratorzy/{id} Administrators GetAdministratorByID
// Returns list of administrators
//
// responses:
//  200: administrator
//  400: error
//  500: error

// swagger:response administrator
type administrator struct {
	// All registered administrators
	// in: body
	Body schemas.Administrator
}

// swagger:parameters CreateAdministrator
type createAdministatorParams struct {
	// Administator to be registered
	// in: body
	Body struct {
		Haslo                       string `json:"haslo"`
		Imie                        string `json:"imie"`
		Login                       string `json:"login"`
		Nazwisko                    string `json:"nazwisko"`
		Pesel                       string `json:"pesel"`
		StanowiskoAdministracyjneId int    `json:"stanowisko_administracyjne_id"`
	}
}

// swagger:route POST /administratorzy/ Administrators CreateAdministrator
// Creates a new administrator
// responses:
//  200: administrator
//  500: error
