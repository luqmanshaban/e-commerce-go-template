package services

import (
	"context"
	"fmt"

	"github.com/luqmanshaban/go-eccomerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductService struct {
	collection *mongo.Collection
}


func NewMongoProductService(collection *mongo.Collection) *MongoProductService {
	return &MongoProductService{
		collection: collection,
	}
}

// Create a new Product
func (s *MongoProductService) CreateProduct(product *models.Product) error {
	// Generate a unique ID for the user
	product.ID = primitive.NewObjectID().Hex()
	_, err := s.collection.InsertOne(context.Background(), product)
	return err
}

// Create Bulk Products
func (s *MongoProductService) CreateProducts(products []*models.Product) ([]interface{}, error) {
	// Create an empty slice to store inserted product IDs
	insertedIDs := make([]interface{}, 0, len(products))

	// Loop through each product in the array
	for _, product := range products {
		// Generate a new ObjectID for each product
		product.ID = primitive.NewObjectID().Hex()

		// Insert the product into the collection
		result, err := s.collection.InsertOne(context.Background(), product)
		if err != nil {
			// Handle error - log the error and potentially return early
			return nil, err
		}

		// If successful, append the inserted product ID to the slice
		insertedIDs = append(insertedIDs, result.InsertedID)
	}

	// Return the slice of inserted product IDs and nil error (if successful)
	return insertedIDs, nil
}

// Get Product By Title
func (s *MongoProductService) GetProductByTitle(title string) (*models.Product, error) {
	var product models.Product
	err := s.collection.FindOne(context.Background(), bson.M{"title": product.Title}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Get Product
func (s *MongoProductService) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	err := s.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Update product
func (s *MongoProductService) UpdateProduct(product *models.Product) error {
	_, err := s.collection.ReplaceOne(context.Background(), bson.M{"_id": product.ID}, product)
	return err
}

// GetAllProducts retrieves all products from the database
func (s *MongoProductService) GetAllProducts() ([]models.Product, error) {
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Delete Product
func (s *MongoProductService) DeleteProduct(id string) error {

	// Define the filter to find the product by its ID
    filter := bson.M{"_id": id}

    // Delete the product using DeleteOne
    result, err := s.collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return err
    }

    // Check if the product was found and deleted
    if result.DeletedCount == 0 {
        return fmt.Errorf("no product found with ID %s", id)
    }

    return nil
}
