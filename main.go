package main

import (
	"api/controllers"
	"api/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	// Register
	router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
	// Login
	router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")
	// Add new task using auth
	router.HandleFunc("/tasks/new", controllers.CreateTask).Methods("POST")
	// Get all tasks using auth
	router.HandleFunc("/me/tasks", controllers.GetTasks).Methods("GET")
	// Change completed status of the task by id
	router.HandleFunc("/me/complete", controllers.UpdateTask).Methods("POST")
	// Delete task by id
	router.HandleFunc("/me/delete", controllers.DeleteTask).Methods("POST")

	// Use JWT auth middleware
	router.Use(middleware.JwtAuthentication)

	// Set port
	port := "8001"

	fmt.Println(port)

	// Start listening
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	} // Catch errors
}