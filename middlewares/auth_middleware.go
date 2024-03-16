package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JwtMiddleware struct {
	secretKey string
}

func NewJwtMiddleware(secretKey string) *JwtMiddleware {
	return &JwtMiddleware{secretKey}
}

func (m *JwtMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "UNAUTHAURIZED", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the token
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
		}

		next.ServeHTTP(w, r)

	})
}

// Helper function to extract token from request header
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
