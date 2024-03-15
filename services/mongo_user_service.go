package services

import (
	"context"
	"github.com/luqmanshaban/go-eccomerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *MongoUserService) CreateUser(user *models.User) error {
	// Generate a unique ID for the user
	user.ID = primitive.NewObjectID().Hex()
	_, err := s.collection.InsertOne(context.Background(), user)
	return err
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
