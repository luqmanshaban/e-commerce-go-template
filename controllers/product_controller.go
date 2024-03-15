package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luqmanshaban/go-eccomerce/models"
	"github.com/luqmanshaban/go-eccomerce/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct {
	productService services.ProductServices
}

func NewProductController(productService services.ProductServices) *ProductController {
	return &ProductController{ productService: productService }
}


// Create a Single Product
func (pc *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	// Check if Product exists
	existingProduct, err := pc.productService.GetProductByTitle(product.Title)
	fmt.Println(existingProduct)
	if err != nil && err != mongo.ErrNoDocuments {
		http.Error(w, "Failed to check for duplicate email", http.StatusInternalServerError)
        return
	}

	if existingProduct != nil {
		http.Error(w, "Duplicate Product", http.StatusBadRequest)
		return
	}


	// Save the product
	err = pc.productService.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
        http.Error(w, "FAILED TO CREATE PRODUCT", http.StatusInternalServerError)
		// println("Failed to create user")
        return
    }

	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// Create Bulk Products
func (pc *ProductController) CreateProducts(w http.ResponseWriter, r *http.Request) {
    var products []*models.Product // Slice to store array of products

    // Decode the request body into the products slice
	fmt.Println(products)
    err := json.NewDecoder(r.Body).Decode(&products)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call productService to create products
    insertedIDs, err := pc.productService.CreateProducts(products)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Handle successful insertion (optional)
    fmt.Println("Successfully created products with IDs:", insertedIDs)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Products created successfully"}) // Example response
}

// Get Product
func (pc *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	user, err := pc.productService.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update Product By Id
func (ac *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.ID = productID
	err = ac.productService.UpdateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Delete Product

func (ac *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productID := params["id"]

	err := ac.productService.DeleteProduct(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := map[string]string{"success": "Product deleted successfully"}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(success)
}