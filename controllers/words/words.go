package words

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/services"
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
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
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
func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {
	var newWord entities.WordRequest
	json.NewDecoder(r.Body).Decode(&newWord)

	ctx := r.Context()

	ownerID, _ := strconv.ParseUint(r.Header.Get("payload_id"), 10, 64)
	newWord.UserID = ownerID

	ctx, tr := tracer.Span(ctx, "controllers.words.create")
	defer tr.End()

	wordCreated, err := ctr.srv.Word.Create(ctx, newWord)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create word: ", newWord)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wordCreated)
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
func (ctr *controllers) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	wordID, _ := strconv.ParseUint(id, 10, 64)
	ownerID, _ := strconv.ParseUint(r.Header.Get("payload_id"), 10, 64)

	ctx, tr := tracer.Span(ctx, "controllers.words.get_by_id")
	defer tr.End()

	word, err := ctr.srv.Word.GetByID(ctx, wordID, ownerID)
	if err != nil {
		ctr.log.ErrorContext(ctx, "Ctrl.GetByid: ", "Error get word by id: ", wordID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(word)
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
func (ctr *controllers) GetAll(w http.ResponseWriter, r *http.Request) {

	ownerID, _ := strconv.ParseUint(r.Header.Get("payload_id"), 10, 64)

	ctx := r.Context()

	ctx, tr := tracer.Span(ctx, "controllers.words.get_all")
	defer tr.End()

	words, err := ctr.srv.Word.GetAll(ctx, ownerID)
	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all word")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(words)
}
