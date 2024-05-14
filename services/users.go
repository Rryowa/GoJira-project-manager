package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rryowa/Gojira-project-manager/entity"
	"github.com/rryowa/Gojira-project-manager/repo"
	"github.com/rryowa/Gojira-project-manager/utils"
)

type UserService struct {
	repo repo.Repo
}

func NewUserService(r repo.Repo) *UserService {
	return &UserService{repo: r}
}

func (us *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", us.HandleUserRegister).Methods("POST")
}

func (us *UserService) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *entity.User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateUserPayload(payload); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,
			utils.ErrorResponse{Error: "Error creating hash"})
		return
	}
	payload.Password = hashedPassword

	//us содержит все поля, обсер происходит в insert
	u, err := us.repo.CreateUser(payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,
			utils.ErrorResponse{Error: "User with this email already exists"})
		fmt.Println(err)
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,
			utils.ErrorResponse{Error: "Error creating cookie"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, token)
}

func validateUserPayload(user *entity.User) error {
	if user.Email == "" {
		return utils.ErrEmailRequired
	}
	if user.Password == "" {
		return utils.ErrPasswordRequired
	}
	return nil
}

func createAndSetAuthCookie(userID int64, w http.ResponseWriter) (string, error) {
	secret := []byte(utils.Envs.JWTSecret)
	token, err := utils.CreateJWT(secret, userID)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
