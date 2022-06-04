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
		// required: true
		// unique: true
		Login string `json:"login"`
		// Haslo has to be longer than 6 characters
		// required: true
		Haslo string `json:"haslo"`
		// required: true
		Imie string `json:"imie"`
		// required: true
		Nazwisko string `json:"nazwisko"`
		// Pesel of the specific administrator. Must be unique and have 11 numbers to be correct.
		// unique: true
		// required: true
		Pesel string `json:"pesel"`
		// Id of administrator's stanowisko, a negative value indicates that he doesn't have one
		// required: true
		StanowiskoAdministracyjneId int `json:"stanowisko_administracyjne_id"`
	}
}

// swagger:route POST /administratorzy/ Administrators CreateAdministrator
// Creates a new administrator
// responses:
//  200: administrator
//  500: error
