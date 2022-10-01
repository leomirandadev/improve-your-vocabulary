package meanings

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
	Create(httpCtx httpRouter.Context)
	GetByID(httpCtx httpRouter.Context)
	GetAll(httpCtx httpRouter.Context)
}

func New(srv *services.Container, log logger.Logger, tokenHasher token.TokenHash) IController {
	return &controllers{srv: srv, log: log, token: tokenHasher}
}

// meaning swagger document
// @Description Create one meaning
// @Tags meaning
// @Param meaning body entities.MeaningRequest true "add meaning"
// @Accept json
// @Produce json
// @Success 201 {object} entities.Meaning
// @Failure 500
// @Security ApiKeyAuth
// @Router /meanings [post]
func (ctr *controllers) Create(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.meanings.create")
	defer tr.End()

	var newMeaning entities.MeaningRequest
	httpCtx.Decode(&newMeaning)

	meaningCreated, err := ctr.srv.Meaning.Create(ctx, newMeaning)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create meaning: ", newMeaning)
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusCreated, meaningCreated)
}

// meaning swagger document
// @Description Get one meaning
// @Tags meaning
// @Param id path string true "Meaning ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Meaning
// @Failure 500
// @Security ApiKeyAuth
// @Router /meanings/{id} [get]
func (ctr *controllers) GetByID(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.meanings.get_by_id")
	defer tr.End()

	id := httpCtx.GetParam("id")
	idMeaning, _ := strconv.ParseUint(id, 10, 64)

	meaning, err := ctr.srv.Meaning.GetByID(ctx, idMeaning)
	if err != nil {
		ctr.log.Error("Ctrl.GetByID: ", "Error get meaning by id: ", idMeaning)
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusOK, meaning)
}

// meaning swagger document
// @Description Get all meaning
// @Tags meaning
// @Accept json
// @Produce json
// @Success 200 {object} []entities.Meaning
// @Failure 500
// @Security ApiKeyAuth
// @Router /meanings [get]
func (ctr *controllers) GetAll(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.meanings.get_all")
	defer tr.End()

	meanings, err := ctr.srv.Meaning.GetAll(ctx)
	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all meaning")
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusOK, meanings)
}
