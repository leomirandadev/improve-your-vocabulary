package users

import (
	"net/http"
	"strconv"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/services"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer"
)

type controllers struct {
	srv   *services.Container
	log   logger.Logger
	token token.TokenHash
}

type IController interface {
	Create(c httpRouter.Context)
	Auth(c httpRouter.Context)
	GetByID(c httpRouter.Context)
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
func (ctr *controllers) Create(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.users.create")
	defer tr.End()

	var newUser entities.UserRequest
	c.Decode(&newUser)

	if err := ctr.srv.User.Create(ctx, newUser); err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create user: ", newUser)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, newUser)
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
func (ctr *controllers) Auth(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.users.auth")
	defer tr.End()

	var userLogin entities.UserAuth
	c.Decode(&userLogin)

	userFound, err := ctr.srv.User.GetUserByLogin(ctx, userLogin)
	if err != nil {
		ctr.log.Error("Ctrl.Auth: ", "Error on find a user", userLogin)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	token, err := ctr.token.Encrypt(userFound)
	if err != nil {
		ctr.log.Error("Ctrl.Auth: ", "Error on generate token", userLogin)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, entities.AuthToken{Token: token})
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
func (ctr *controllers) GetByID(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.users.get_by_id")
	defer tr.End()

	idUser, _ := strconv.ParseUint(c.GetParam("id"), 10, 64)

	user, err := ctr.srv.User.GetByID(ctx, idUser)
	if err != nil {
		ctr.log.Error("Ctrl.GetByid: ", "Error get user by id: ", idUser)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}
