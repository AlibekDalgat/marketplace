package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	loginCtx            = "login"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "пустой хедер аутентификации")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "неправильный header")
		return
	}

	login, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(loginCtx, login)
}

func getUserLogin(c *gin.Context) (string, error) {
	login, ok := c.Get(loginCtx)
	if !ok {
		return "", errors.New("не найден логин пользователя")
	}
	loginStr, ok := login.(string)
	if !ok {
		return "", errors.New("неверный тип логина пользователя")
	}

	return loginStr, nil
}
