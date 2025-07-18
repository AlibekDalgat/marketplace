package delivery

import (
	"github.com/gin-gonic/gin"
	"marketplace/internal/models"
	"math"
	"net/http"
	"strconv"
)

func (h *Handler) createAdvert(c *gin.Context) {
	login, err := getUserLogin(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if login == "" {
		newErrorResponse(c, http.StatusUnauthorized, "неавторизированный пользователь")
		return
	}

	var input models.Advert
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	output, err := h.services.Create(login, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}

type getAllAdvertsResponse struct {
	Data []models.AdvertOutput `json:"data"`
}

func (h *Handler) listAdverts(c *gin.Context) {
	login, _ := getUserLogin(c)
	var params models.AdvertParams
	params.Sort = c.DefaultQuery("sort", "posting_date")
	params.Direction = c.DefaultQuery("direction", "desc")

	var err error
	params.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный параметр limit")
		return
	}
	params.Page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный параметр page")
		return
	}
	params.PriceMin, err = strconv.Atoi(c.DefaultQuery("pricemin", "0"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный параметр pricemin")
		return
	}
	params.PriceMax, err = strconv.Atoi(c.DefaultQuery("pricemax", strconv.Itoa(math.MaxInt)))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверный параметр pricemax")
		return
	}

	adverts, err := h.services.GetAll(login, params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllAdvertsResponse{Data: adverts})
}
