package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	// "github.com/luqmanshaban/go-eccomerce/initializers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	// initializers.LoadEnv()
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the authorization header
        authHeader := r.Header.Get("Authorization")

        // Check if the authorization header is empty
        if authHeader == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Extract the token from the authorization header
        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            http.Error(w, "Invalid token format", http.StatusBadRequest)
            return
        }
        tokenString := tokenParts[1]

        // Verify the token
        claims := &jwt.StandardClaims{}
		jwt_secret := os.Getenv("JWT_KEY")
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwt_secret), nil // Replace "your-secret-key" with your actual secret key
        })
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Check if the token is valid
        if !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Token is valid, call the next handler
        next.ServeHTTP(w, r)
    })
}
