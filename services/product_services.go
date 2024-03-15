package services

import "github.com/luqmanshaban/go-eccomerce/models"

type ProductServices interface {
    CreateProduct(product *models.Product) error
    CreateProducts([]*models.Product) ([]interface{}, error)
    GetProductByTitle(title string) (*models.Product, error)
    GetProductByID(id string) (*models.Product, error)
    GetAllProducts() ([]models.Product, error) // New method for getting all products
    UpdateProduct(product *models.Product) error
    DeleteProduct(id string) error
}
