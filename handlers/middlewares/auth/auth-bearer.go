package auth

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
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

func (m *middlewareJWT) Public(next httpRouter.HandlerFunc) httpRouter.HandlerFunc {

	return func(httpCtx httpRouter.Context) {

		if err := m.verifyRoles(httpCtx.Headers(), false); err != nil {
			httpCtx.JSON(http.StatusUnauthorized, Response{Message: "Permissão negada"})
			return
		}

		next(httpCtx)
	}

}

func (m *middlewareJWT) Private(next httpRouter.HandlerFunc) httpRouter.HandlerFunc {

	return func(httpCtx httpRouter.Context) {

		if err := m.verifyRoles(httpCtx.Headers(), true, entities.Roles...); err != nil {
			httpCtx.JSON(http.StatusUnauthorized, Response{Message: "Permissão negada"})
			return
		}

		next(httpCtx)
	}

}

func (m *middlewareJWT) Admin(next httpRouter.HandlerFunc) httpRouter.HandlerFunc {

	return func(httpCtx httpRouter.Context) {

		if err := m.verifyRoles(httpCtx.Headers(), true, "admin"); err != nil {
			httpCtx.JSON(http.StatusUnauthorized, Response{Message: "Permissão negada"})
			return
		}

		next(httpCtx)
	}

}

func (m *middlewareJWT) verifyRoles(header http.Header, logged bool, roles ...string) error {

	if !logged {
		return nil
	}

	if header["Authorization"] == nil {
		m.log.Error("authorization nil")
		return errors.New("WITHOUT_AUTHORIZATION")
	}

	bearerSplited := strings.Split(header["Authorization"][0], " ")
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
			m.InsertTokenFieldsOnPayload(claims, header)
			return nil
		}
	}

	m.log.Error("role not found")
	return errors.New("UNAUTHORIZED")
}

func (m *middlewareJWT) InsertTokenFieldsOnPayload(token map[string]interface{}, header http.Header) {
	header.Add("payload_id", strconv.FormatInt(int64(token["id"].(float64)), 10))
	header.Add("payload_name", token["name"].(string))
	header.Add("payload_nick_name", token["nick_name"].(string))
	header.Add("payload_email", token["email"].(string))
	header.Add("payload_role", token["role"].(string))
}
