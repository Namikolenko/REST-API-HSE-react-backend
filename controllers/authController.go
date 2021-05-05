package controllers

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
)

func CreateAccount(w http.ResponseWriter, r *http.Request)  {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := user.Create()
	utils.Respond(w, resp)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Login, user.Password)
	utils.Respond(w, resp)
}
