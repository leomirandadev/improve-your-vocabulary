package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
)

type middlewareJWT struct {
	token token.TokenHash
	log   logger.Logger
}

type Response struct {
	Message string `json:"message"`
}

func NewBearer(tokenHasher token.TokenHash, log logger.Logger) AuthMiddleware {
	return &middlewareJWT{
		token: tokenHasher,
		log:   log,
	}
}

func (m *middlewareJWT) Public(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := m.VerifyRoles(r, false); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Response{Message: "Permissão negada"})
			return
		}

		next.ServeHTTP(w, r)
	})

}

func (m *middlewareJWT) Admin(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := m.VerifyRoles(r, true, "admin"); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(Response{Message: "Permissão negada"})
			return
		}

		next.ServeHTTP(w, r)
	})

}

func (m *middlewareJWT) VerifyRoles(r *http.Request, logged bool, roles ...string) error {

	if !logged {
		return nil
	}

	if r.Header["Authorization"] == nil {
		m.log.Error("authorization nil")
		return errors.New("WITHOUT_AUTHORIZATION")
	}

	bearerSplited := strings.Split(r.Header["Authorization"][0], " ")
	if len(bearerSplited) != 2 {
		m.log.Error("can't split bearer")
		return errors.New("INVALID_AUTHORIZATION")
	}

	isValid, claims, err := m.token.Decrypt(bearerSplited[1])
	if err != nil {
		m.log.Error("decrypt", err)
		return err
	}

	if !isValid {
		m.log.Error("TOKEN NOT VALID")
		return errors.New("INVALID_AUTHORIZATION")
	}

	for _, role := range roles {
		if claims["role"] == role {
			m.InsertTokenFieldsOnPayload(claims, r)
			return nil
		}
	}

	m.log.Error("role not found")
	return errors.New("UNAUTHORIZED")
}

func (m *middlewareJWT) InsertTokenFieldsOnPayload(token map[string]interface{}, r *http.Request) {
	r.Header.Add("payload_id", strconv.FormatInt(int64(token["id"].(float64)), 10))
	r.Header.Add("payload_name", token["name"].(string))
	r.Header.Add("payload_nick_name", token["nick_name"].(string))
	r.Header.Add("payload_email", token["email"].(string))
	r.Header.Add("payload_role", token["role"].(string))
}
