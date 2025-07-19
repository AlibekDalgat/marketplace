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

func (h *Handler) userIdentity(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			if required {
				newErrorResponse(c, http.StatusUnauthorized, "пустой хедер аутентификации")
				return
			}
			c.Next()
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			if required {
				newErrorResponse(c, http.StatusUnauthorized, "неправильный header")
				return
			}
			c.Next()
			return
		}

		login, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			if required {
				newErrorResponse(c, http.StatusUnauthorized, err.Error())
				return
			}
			c.Next()
			return
		}

		c.Set(loginCtx, login)
		c.Next()
	}
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
