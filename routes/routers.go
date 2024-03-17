// routes/routes.go

package routes

import (
	"github.com/gorilla/mux"
	"github.com/luqmanshaban/go-eccomerce/controllers"
	"github.com/luqmanshaban/go-eccomerce/initializers"
	"github.com/luqmanshaban/go-eccomerce/middlewares"
	"github.com/luqmanshaban/go-eccomerce/services"
)

func SetupRoutes(router *mux.Router, userService services.UserService, productService services.ProductServices) {
    initializers.LoadEnv()
    // Create instances of controllers with injected dependencies
    authController := controllers.NewAuthController(userService)
    productController := controllers.NewProductController(productService)

    // Routes for auth controller
    router.HandleFunc("/api/register", authController.CreateUser).Methods("POST")
    router.HandleFunc("/api/login", authController.Login).Methods("POST")
    router.HandleFunc("/api/users/{id}", authController.GetUserByID).Methods("GET")
    router.HandleFunc("/api/users/{id}", authController.UpdateUser).Methods("PUT")
    router.HandleFunc("/api/users/{id}", authController.DeleteUser).Methods("DELETE")

    // Routes for Product controller
    router.HandleFunc("/api/products", productController.CreateProduct).Methods("POST")
    router.HandleFunc("/api/products/bulk", productController.CreateProducts).Methods("POST")
    router.HandleFunc("/api/products/bulk", productController.GetAllProducts).Methods("GET")
    router.HandleFunc("/api/products/{id}", productController.GetProductByID).Methods("GET")

    // Apply middleware to PUT, DELETE requests on product routes
    productRoutes := router.PathPrefix("/api/products").Subrouter()
    productRoutes.Use(middlewares.AuthMiddleware)
    productRoutes.HandleFunc("/{id}", productController.UpdateProduct).Methods("PUT")
    productRoutes.HandleFunc("/{id}", productController.DeleteProduct).Methods("DELETE")
}
