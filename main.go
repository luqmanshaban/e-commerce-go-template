package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
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

     // CORS middleware
     corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Start the HTTP server
    fmt.Println("SERVER RUNNING ON http://localhost:3000")
    http.ListenAndServe(":4000", corsHandler(router))
}
