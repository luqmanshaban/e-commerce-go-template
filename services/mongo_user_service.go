package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	// "fmt"

	"github.com/luqmanshaban/go-eccomerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type MongoUserService struct {
	collection *mongo.Collection
}

// GetUserByEmail implements UserService.
func (s *MongoUserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByUsername implements UserService.
func (s *MongoUserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := s.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// DeleteUser implements UserService.
func (s *MongoUserService) DeleteUser(id string) error {
	panic("unimplemented")
}

func NewMongoUserService(collection *mongo.Collection) *MongoUserService {
	return &MongoUserService{
		collection: collection,
	}
}

func (s *MongoUserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	// Assuming user.ID is set somewhere before calling this method
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserService) UpdateUser(user *models.User) error {
	_, err := s.collection.ReplaceOne(context.Background(), bson.M{"_id": user.ID}, user)
	return err
}

func (s *MongoUserService) VerifyEmail(verificationCode string) error {

	println(verificationCode)
    var user models.User
    err := s.collection.FindOne(context.Background(), bson.M{"verification_code": verificationCode}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return errors.New("invalid verification code") // User not found or invalid code
        }
        return err
    }

    // Parse the verification code expiration time
    expirationTime, err := time.Parse(time.RFC3339, user.VerificationCodeExpiration)
    if err != nil {
        return err // Handle parsing error
    }

    // Check if verification code is expired
    if time.Now().After(expirationTime) {
        return errors.New("verification code has expired")
    }

    // Update user to verified status
    user.IsVerified = true                    // Set boolean flag to mark as verified
    user.VerificationCode = ""                // Clear verification code
    user.VerificationCodeExpiration = ""      // Clear expiration time

    _, err = s.collection.ReplaceOne(context.Background(), bson.M{"_id": user.ID}, user)
    if err != nil {
        return err
    }

    return nil // Verification successful
}



  func (s *MongoUserService) CreateUser(user *models.User) error {
	// Hash the password before storing it
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hash)
	fmt.Println("Hashed password:", user.Password) // Add this line

	// Generate a unique ID for the user
	user.ID = primitive.NewObjectID().Hex()
	_, err = s.collection.InsertOne(context.Background(), user)
	return err
}

// Login implements UserService.
func (s *MongoUserService) Login(email, password string) (*models.User, error) {
    // Get user by email
    user, err := s.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }

    // Check if user exists
    if user == nil {
        return nil, errors.New("user not found")
    }

    // Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(string(user.Password)), []byte(string(password)))
    if err != nil {
        return nil, errors.New(err.Error())
    }

    return user, nil
	
}
