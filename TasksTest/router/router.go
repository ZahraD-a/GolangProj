package router

import (
	"hello/tasks/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/task", controller.CreateTasks).Methods("POST")      //CREATE
	router.HandleFunc("/Alltasks", controller.GetMyAlltasks).Methods("GET") // READ
	//	router.HandleFunc("/update/{id}", controller.UpdateTasks).Methods("PUT")       // UPDATE
	router.HandleFunc("/deleteOne/{id}", controller.DeleteAtask).Methods("DELETE") //DELETE ONE
	router.HandleFunc("/deleteAll", controller.DeleteAlltasks).Methods("DELETE")   // DELETE ALL

	return router
}
