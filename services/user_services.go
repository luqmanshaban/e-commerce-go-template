package services

import "github.com/luqmanshaban/go-eccomerce/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}
