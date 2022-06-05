package routes

// here we define all the routes

import (
	"crud/pkg/controller"
	"crud/pkg/middleware"
	"crud/pkg/signin"
	"github.com/gorilla/mux"
)

var Userroutes = func(router *mux.Router) {
	router.HandleFunc("/create/", middleware.Authorize(controller.Createuser)).Methods("POST")
	router.HandleFunc("/get/{offset}/{limit}", middleware.Authorize(controller.Getuser)).Methods("GET")
	router.HandleFunc("/getbyid{id}/", middleware.Authorize(controller.Getbyid)).Methods("GET")
	router.HandleFunc("/delete{id}/", middleware.Authorize(controller.Deleteuser)).Methods("DELETE")
	router.HandleFunc("/update{id}/", middleware.Authorize(controller.Updateuser)).Methods("PUT")
	router.HandleFunc("/login", login.Login).Methods("POST")

}
