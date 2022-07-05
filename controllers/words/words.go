package words

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newWord entities.WordRequest
	json.NewDecoder(r.Body).Decode(&newWord)

	ctx := r.Context()

	wordCreated, err := ctr.srv.Word.Create(ctx, newWord)
	if err != nil {
		ctr.log.Error("Ctrl.Create: ", "Error on create word: ", newWord)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wordCreated)
}

func (ctr *controllers) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idWord, _ := strconv.ParseUint(params["id"], 10, 64)

	ctx := r.Context()
	word, err := ctr.srv.Word.GetByID(ctx, idWord)

	if err != nil {
		ctr.log.Error("Ctrl.GetByid: ", "Error get word by id: ", idWord)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(word)
}

func (ctr *controllers) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	words, err := ctr.srv.Word.GetAll(ctx)

	if err != nil {
		ctr.log.Error("Ctrl.GetAll: ", "Error get all word")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(words)
}
