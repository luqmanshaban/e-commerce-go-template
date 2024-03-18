## Go E-commerce Server Documentation

This document provides an overview of a Go e-commerce server project structure, functionalities, and current development stage.

### Project Structure

The project is organized into the following directories:

* **config:** Stores configuration files, including database connection details (`db.go`).
* **controllers:** Houses controllers for handling API requests (`controllers`).
* **example.env:** Provides a template for environment variables (`.env`).
* **go.mod:** Declares project dependencies (`go.mod`).
* **go.sum:** Contains checksums for dependencies (`go.sum`).
* **initializers:** Code for initialization tasks (currently empty, `initializers`).
* **main.go:** The main entry point of the application.
* **middlewares:** Contains middleware functions for request processing (`middlewares`).
* **models:** Defines data structures representing entities like users and products (`models`).
* **routes:** Handles setting up routes for different API endpoints (`routes`).
* **services:** Business logic layer for interacting with data and functionality (`services`).
* **vendor:** Stores downloaded dependencies (`vendor`).

### Functionalities

The current functionalities of the server include:

* User Management:
    * User registration (`/api/register`).
    * User login (`/api/login`).
    * Get user by ID (protected, `/api/users/{id}`).
    * Update user details (protected, `/api/users/{id}`).
    * Delete user (protected, `/api/users/{id}`).
    * User email verification (`/api/users/verify-email/{code}`).
* Product Management (basic):
    * Create a product (`/api/products`).
    * Create multiple products in bulk (`/api/products/bulk`).
    * Retrieve all products (for now bulk created products, `/api/products/bulk`).
    * Get product by ID (`/api/products/{id}`).

**Note:** Update and Delete functionalities for products are implemented and functional; however, they require authentication middleware for security. The current implementation provides basic functionality and can be enhanced for more robust authentication and authorization mechanisms.

### Authentication

The server implements basic authentication using JWT (middleware not fully implemented yet). Users can register and login to obtain access tokens required for protected routes (update/delete functionalities).

### Database

The server uses MongoDB for data storage. The connection and configuration are handled in the `config` directory.

### Current Development Stage

The project is under development. While core functionalities like user registration, login, basic product management, and email verification are functional, some features require further implementation:

* User functionalities beyond basic CRUD (update, read, delete) like searching or filtering.
* Full product management with update and delete functionalities with proper authentication.
* Shopping cart and order processing functionalities (models like `cart.go` and `order.go` are present but not implemented).
* Payment processing integration (model `payment.go` exists but not implemented).

### Dependencies

The project relies on various Go packages:

* **gorilla/mux:** Routing library for handling API requests.
* **gorilla/handlers:** Middleware for handling CORS (Cross-Origin Resource Sharing).
* **go.mongodb.org/mongo-driver:** Official MongoDB driver for Go.
* **golang.org/x/crypto/bcrypt:** For password hashing.

Environment variables like database connection string and JWT secret key are stored in a separate `.env` file (not included in this documentation).

This documentation provides a starting point for understanding the project structure, functionalities, and current development stage. Please refer to the actual code for detailed implementation and future updates.


### License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
