package routes

import (
	"users_manager/internal/domain/controllers"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")

	router.HandleFunc("/getUserById/{id}", controllers.GetUserById).Methods("GET")

	router.HandleFunc("/getAllUsers", controllers.GetAllUsers).Methods("GET")

	router.HandleFunc("/createUser", controllers.CreateUser).Methods("POST")

	router.HandleFunc("/updateUser", controllers.UpdateUser).Methods("PUT")

	router.HandleFunc("/deleteUser/{id}", controllers.DeleteUserById).Methods("DELETE")

}
