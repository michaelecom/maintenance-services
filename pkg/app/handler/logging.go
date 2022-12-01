package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rimeks.ru/services/pkg/app/structures"
)

func (h *Handler) createLog(c *gin.Context) {
	var input structures.LogInput
	c.Set("input", &input)

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверный формат переданных данных")
		return
	}

	id, err := h.services.Logging.CreateLog(input)
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке добавления записи о событии")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ids": id})
}

func (h *Handler) getAllLogs(c *gin.Context) {
	orders, err := h.services.Logging.GetAllLogs()
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке получения журнала событий")
		return
	}

	c.IndentedJSON(http.StatusOK, orders)
}

func (h *Handler) getAllLogsByServiceMarketID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("service_market_id"))
	if err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверно указан ID сервис - маркета")
		return
	}

	orders, err := h.services.Logging.GetAllLogsByServiceMarketID(id)
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке получения журнала событий по ID сервис - маркета")
		return
	}

	c.IndentedJSON(http.StatusOK, orders)
}

func (h *Handler) clearLogs(c *gin.Context) {
	if err := h.services.Logging.ClearLogs(); err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при выполнении метода очистки")
		return
	}

	c.Status(http.StatusOK)
}
