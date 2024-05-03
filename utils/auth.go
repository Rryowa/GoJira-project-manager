package utils

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rryowa/Gojira-project-manager/repo"
	"golang.org/x/crypto/bcrypt"
)

// Wrap handlers in authentication

func WithJWTAuth(handlerFunc http.HandlerFunc, rp repo.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Get token from request
		tokenString := GetTokenFromRequest(r)
		//Validate
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			PermissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token", w)
			PermissionDenied(w)
			return
		}
		//Get claims for token
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)

		//Get Id from token(request)
		_, err = rp.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			PermissionDenied(w)
			return
		}

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	// api/v1/url?query=token
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		log.Println(tokenAuth)
		return tokenAuth
	}
	if tokenQuery != "" {
		log.Println(tokenQuery)
		return tokenQuery
	}
	//if token is empty
	return ""
}

func validateJWT(t string) (*jwt.Token, error) {
	secret := NewConfig().JWTSecret
	return jwt.Parse(t, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CreateJWT(secret []byte, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(time.Hour * 24 * 120).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
