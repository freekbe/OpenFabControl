package handler

import (
	"OpenFabControl/database"
	"OpenFabControl/model"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Login(w http.ResponseWriter, r* http.Request) {

	if reject_all_methode_exept(r, w, http.MethodPost) != nil { return }

	var payload struct {
		EMAIL string `json:"email"`
		PASSWORD string `json:"password"`
	}

	if extract_payload_data(r, w, &payload) != nil { return }

	if !validate_payload(payload.EMAIL == "", "email cannot be empty", w) { return }
	if !validate_payload(payload.PASSWORD == "", "password cannot be empty", w) { return }

	// check password
	var hash, status string;
	var user_id int;
	err := database.Self.QueryRow(`SELECT password, id, status FROM users WHERE email = $1`, payload.EMAIL).Scan(&hash, &user_id, &status)
	if err != nil {
		if err == sql.ErrNoRows{
			http.Error(w, "Invalid credential", http.StatusForbidden)
			return
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	if !checkPasswordHash(payload.PASSWORD, hash) {
		http.Error(w, "Invalid credential", http.StatusForbidden)
		return
	}

	// JWT token
	expiration_time := time.Now().Add(24*time.Hour)
	claims := model.Claims {
		USERID: user_id,
		EMAIL: payload.EMAIL,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration_time),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_TOKEN"))
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"token": tokenString})
}
