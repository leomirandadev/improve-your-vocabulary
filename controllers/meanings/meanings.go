package meanings

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
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
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
func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {
	var newMeaning entities.MeaningRequest
	json.NewDecoder(r.Body).Decode(&newMeaning)

	ctx := r.Context()

	meaningCreated, err := ctr.srv.Meaning.Create(ctx, newMeaning)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create meaning: ", newMeaning)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(meaningCreated)
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
func (ctr *controllers) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idMeaning, _ := strconv.ParseUint(id, 10, 64)

	ctx := r.Context()

	meaning, err := ctr.srv.Meaning.GetByID(ctx, idMeaning)
	if err != nil {
		ctr.log.Error("Ctrl.GetByID: ", "Error get meaning by id: ", idMeaning)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meaning)
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
func (ctr *controllers) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	meanings, err := ctr.srv.Meaning.GetAll(ctx)
	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all meaning")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meanings)
}
