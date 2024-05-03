package utils

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var ErrNameRequired = errors.New("name is required")
var ErrProjectIDRequired = errors.New("project id is required")
var ErrUserIDRequired = errors.New("user id is required")

var ErrEmailRequired = errors.New("email is required")
var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("last name is required")
var ErrPasswordRequired = errors.New("password is required")

func PermissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "Permission denied"})
}
