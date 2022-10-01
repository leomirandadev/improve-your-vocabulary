package words

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

// word swagger document
// @Description Create one word
// @Tags word
// @Param word body entities.WordRequest true "add word"
// @Accept json
// @Produce json
// @Success 201 {object} entities.Word
// @Failure 500
// @Security ApiKeyAuth
// @Router /words [post]
func (ctr *controllers) Create(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.create")
	defer tr.End()

	var newWord entities.WordRequest
	httpCtx.Decode(&newWord)

	ownerID, _ := strconv.ParseUint(httpCtx.GetParam("payload_id"), 10, 64)
	newWord.UserID = ownerID

	wordCreated, err := ctr.srv.Word.Create(ctx, newWord)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create word: ", newWord)
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusCreated, wordCreated)
}

// word swagger document
// @Description Get one word
// @Tags word
// @Param id path string true "Word ID"
// @Accept json
// @Produce json
// @Success 200 {object} entities.Word
// @Failure 500
// @Security ApiKeyAuth
// @Router /words/{id} [get]
func (ctr *controllers) GetByID(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.get_by_id")
	defer tr.End()

	id := httpCtx.GetParam("id")
	wordID, _ := strconv.ParseUint(id, 10, 64)
	ownerID, _ := strconv.ParseUint(httpCtx.GetParam("payload_id"), 10, 64)

	word, err := ctr.srv.Word.GetByID(ctx, wordID, ownerID)
	if err != nil {
		ctr.log.ErrorContext(ctx, "Ctrl.GetByid: ", "Error get word by id: ", wordID)
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusOK, word)
}

// word swagger document
// @Description Get all word
// @Tags word
// @Accept json
// @Produce json
// @Success 200 {object} []entities.Word
// @Failure 500
// @Security ApiKeyAuth
// @Router /words [get]
func (ctr *controllers) GetAll(httpCtx httpRouter.Context) {
	ctx := httpCtx.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.get_all")
	defer tr.End()

	ownerID, _ := strconv.ParseUint(httpCtx.GetFromHeader("payload_id"), 10, 64)

	words, err := ctr.srv.Word.GetAll(ctx, ownerID)
	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all word")
		httpCtx.JSON(http.StatusInternalServerError, nil)
		return
	}

	httpCtx.JSON(http.StatusOK, words)
}
