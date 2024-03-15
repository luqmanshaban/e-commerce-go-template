package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luqmanshaban/go-eccomerce/models"
	"github.com/luqmanshaban/go-eccomerce/services"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{ userService: userService }
}

func (ac *AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Check if the email already exists in the database
	existingUser, err := ac.userService.GetUserByEmail(user.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, "Failed to check for duplicate email", http.StatusInternalServerError)
        return
	}

	if existingUser != nil {
		http.Error(w, "Duplicate Email", http.StatusBadRequest)
		return
	}

	// Check if the username is exists in the db
	existingUsername, err := ac.userService.GetUserByUsername(user.Username)
	if err != nil  && err != mongo.ErrNoDocuments{
		http.Error(w, "Failed to check for duplicate", http.StatusInternalServerError)
		return
	}

	if existingUsername != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}


    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    err = ac.userService.CreateUser(&user)
    if err != nil {
		fmt.Println(err.Error())
        http.Error(w, "FAILED TO CREATE USER", http.StatusInternalServerError)
		// println("Failed to create user")
        return
    }

    // Clear the password before sending the user object in the response
    user.Password = ""

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}


func (ac *AuthController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := ac.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (ac *AuthController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = userID
	err = ac.userService.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (ac *AuthController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	err := ac.userService.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}