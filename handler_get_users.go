package main

import (
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {

	users, error := apiCfg.DB.GetAllUsers(r.Context())
	if error != nil {
		log.Fatal("Error getting users. Examine the request code.")
	}

	respondWithJSON(w, 200, databaseConvertUserstoUsers(users))
}
