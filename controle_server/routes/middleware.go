package routes

import (
	"OpenFabControl/model"
	"OpenFabControl/utils"
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// middleware to check if a user is authentificated with a JWT token
func auth_middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// remove "Bearer from token"
		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		// validate token
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_TOKEN")), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if utils.Reject_user_status(w, claims.USERID, []string{"pending", "desactivated"}) != nil {
			http.Error(w, "Your account is desactivated or pending activation, if you're part of the system, contact the administrator", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), "user_id", claims.USERID)
		ctx = context.WithValue(ctx, "username", claims.EMAIL)

		next(w, r.WithContext(ctx))
	}
}
