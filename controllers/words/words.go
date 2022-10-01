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
	Create(c httpRouter.Context)
	GetByID(c httpRouter.Context)
	GetAll(c httpRouter.Context)
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
func (ctr *controllers) Create(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.create")
	defer tr.End()

	var newWord entities.WordRequest
	c.Decode(&newWord)

	newWord.UserID, _ = strconv.ParseUint(c.GetParam("payload_id"), 10, 64)

	wordCreated, err := ctr.srv.Word.Create(ctx, newWord)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create word: ", newWord)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, wordCreated)
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
func (ctr *controllers) GetByID(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.get_by_id")
	defer tr.End()

	wordID, _ := strconv.ParseUint(c.GetParam("id"), 10, 64)
	ownerID, _ := strconv.ParseUint(c.GetParam("payload_id"), 10, 64)

	word, err := ctr.srv.Word.GetByID(ctx, wordID, ownerID)
	if err != nil {
		ctr.log.ErrorContext(ctx, "Ctrl.GetByid: ", "Error get word by id: ", wordID)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, word)
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
func (ctr *controllers) GetAll(c httpRouter.Context) {
	ctx := c.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.get_all")
	defer tr.End()

	ownerID, _ := strconv.ParseUint(c.GetFromHeader("payload_id"), 10, 64)

	words, err := ctr.srv.Word.GetAll(ctx, ownerID)
	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all word")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, words)
}
