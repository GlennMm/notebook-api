package user

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func RegisterUserRoute(router *mux.Router) {
	// sub_routes := router.PathPrefix("/user").Subrouter()

	// register sub-route here
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}

// JWT SETUP
var jwtSecret = []byte("hsdhafhasfhsadsadasddgdgasdgasd.gasds,ahsadfassdhasdhadshasdhasdhasd8gasdgasdgadsg5asfghs34fhas56fahfhs")

type Claim struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func generateToken(userID int) (string, error) {
	claims := &Claim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func authenticate(username string, password string) (int, error) {
	return 1, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		userID, err := authenticate(username, password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := generateToken(userID)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"token": "` + token + `"}`))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Get the Authorization header from the request
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization token missing", http.StatusUnauthorized)
		return
	}

	// Remove the "Bearer " prefix from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			http.Error(w, "Invalid token signature", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	// Validate token and extract claims
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		// Now you have access to the user's claims, such as claims.UserID
		userID := claims.UserID

		// Perform authorized actions based on the user's claims
		// For example, you can fetch user data from the database and return it as JSON
		user := fetchUserFromDatabase(userID)
		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		jsonResponse, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	} else {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}
}
