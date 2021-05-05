package controllers

import (
	"api/models"
	u "api/utils"
	"encoding/json"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	task := &models.Task{}

	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	task.UserId = user
	resp := task.Create()
	u.Respond(w, resp)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTasks(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	task := &models.Task{}
	task.ID = 0
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	info := models.ChangeStateById(task.Id, id)
	u.Respond(w, info)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	task := &models.Task{}
	task.ID = 0
	err := json.NewDecoder(r.Body).Decode(task)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	info := models.DeleteTask(task.Id, id)
	u.Respond(w, info)
}
