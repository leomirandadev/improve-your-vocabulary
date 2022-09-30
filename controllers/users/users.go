package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/services"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
)

type controllers struct {
	srv   *services.Container
	log   logger.Logger
	token token.TokenHash
}

type IController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
}

func New(srv *services.Container, log logger.Logger, tokenHasher token.TokenHash) IController {
	return &controllers{srv: srv, log: log, token: tokenHasher}
}

// user swagger document
// @Description Create one user
// @Tags user
// @Param user body entities.UserRequest true "create new user"
// @Accept json
// @Produce json
// @Success 201 {object} entities.UserRequest
// @Failure 500
// @Security ApiKeyAuth
// @Router /users [post]
func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {

	var newUser entities.UserRequest
	json.NewDecoder(r.Body).Decode(&newUser)

	ctx := r.Context()
	err := ctr.srv.User.Create(ctx, newUser)

	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create user: ", newUser)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// user swagger document
// @Description Authenticate user
// @Tags user
// @Param user body entities.UserAuth true "add user"
// @Accept json
// @Produce json
// @Success 200 {object} entities.AuthToken
// @Failure 500
// @Failure 400
// @Security ApiKeyAuth
// @Router /users/auth [post]
func (ctr *controllers) Auth(w http.ResponseWriter, r *http.Request) {

	var userLogin entities.UserAuth
	json.NewDecoder(r.Body).Decode(&userLogin)

	ctx := r.Context()
	userFound, err := ctr.srv.User.GetUserByLogin(ctx, userLogin)

	if err != nil {
		ctr.log.Error("Ctrl.Auth: ", "Error on find a user", userLogin)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := ctr.token.Encrypt(userFound)
	if err != nil {
		ctr.log.Error("Ctrl.Auth: ", "Error on generate token", userLogin)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entities.AuthToken{Token: token})
}

// user swagger document
// @Description Get one user
// @Tags user
// @Param id path string true "User ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.User
// @Failure 500
// @Security ApiKeyAuth
// @Router /users/{id} [get]
func (ctr *controllers) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	idUser, _ := strconv.ParseUint(id, 10, 64)

	user, err := ctr.srv.User.GetByID(ctx, idUser)
	if err != nil {
		ctr.log.Error("Ctrl.GetByid: ", "Error get user by id: ", idUser)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
