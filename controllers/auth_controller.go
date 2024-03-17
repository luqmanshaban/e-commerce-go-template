package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/luqmanshaban/go-eccomerce/models"
	"github.com/luqmanshaban/go-eccomerce/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	userService services.UserService
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{userService: userService}
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
	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, "Failed to check for duplicate", http.StatusInternalServerError)
		return
	}

	if existingUsername != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}


	err = ac.userService.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "FAILED TO CREATE USER", http.StatusInternalServerError)
		// println("Failed to create user")
		return
	}


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

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the request body into loginRequest struct
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform login
	user, err := ac.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiry time
	})

	// Sign and get the complete encoded token as a string
	jwt_secret := os.Getenv("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Clear the password before sending the user object in the response
	user.Password = ""

	// Respond with the user details and access token
	response := struct {
		User        *models.User `json:"user"`
		AccessToken string       `json:"access_token"`
	}{
		User:        user,
		AccessToken: tokenString,
	}

	// Encode response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
