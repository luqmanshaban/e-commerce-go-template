package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/luqmanshaban/go-eccomerce/config"
    "github.com/luqmanshaban/go-eccomerce/routes"
    "github.com/luqmanshaban/go-eccomerce/services"
)

func main() {
    // Connect to the database
    config.ConnectToDB()

    // Create a new *mux.Router instance
    router := mux.NewRouter()

    // Get a reference to the MongoDB collection for users, products
    userCollection := config.DB.Collection("users")
    productCollection := config.DB.Collection("products")

    // Create an instance of MongoUserService with the user, products collection
    userService := services.NewMongoUserService(userCollection)
    productService := services.NewMongoProductService(productCollection)


    // Setup routes by passing *mux.Router and services
    routes.SetupRoutes(router, userService, productService)

    // Start the HTTP server
    fmt.Println("SERVER RUNNING ON http://localhost:3000")
    http.ListenAndServe(":3000", router)
}
