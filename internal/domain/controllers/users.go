package controllers

import (
	"encoding/json"
	"net/http"
	"users_manager/internal/domain/entities"
	usecases "users_manager/internal/domain/usecases/user"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid Param", http.StatusBadRequest)
		return
	}
	response := usecases.GetUser(id)
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response := usecases.GetAllUsers()
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userToCreate entities.User
	json.NewDecoder(r.Body).Decode(&userToCreate)
	usecases.Create(userToCreate)
	json.NewEncoder(w).Encode(userToCreate)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)
	response := usecases.Update(user)
	json.NewEncoder(w).Encode(response)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Invalid Param", http.StatusBadRequest)
		return
	}
	response := usecases.Delete(id)
	json.NewEncoder(w).Encode(response)
}
