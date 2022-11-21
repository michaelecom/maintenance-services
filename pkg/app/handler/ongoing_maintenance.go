package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rimeks.ru/services/pkg/app/structures"
)

func (h *Handler) createOrder(c *gin.Context) {
	var input structures.OrderList
	c.Set("input", &input)

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверный формат переданных данных")
		return
	}

	id, err := h.services.OngoingMaintenance.CreateOrder(input)
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке добавления записи")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders, err := h.services.OngoingMaintenance.GetAllOrders()
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке получения списков обслуживания")
		return
	}

	c.IndentedJSON(http.StatusOK, orders)
}

func (h *Handler) getAllOrdersByServiceMarketID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("service_market_id"))
	if err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверно указан ID сервис - маркета")
		return
	}

	orders, err := h.services.OngoingMaintenance.GetAllOrdersByServiceMarketID(id)
	if err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке получения списка по ID сервис - маркета")
		return
	}

	c.IndentedJSON(http.StatusOK, orders)
}

func (h *Handler) updateOrder(c *gin.Context) {
	var input structures.OrderList
	c.Set("input", &input)

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверный формат переданных данных")
		return
	}

	if err := h.services.OngoingMaintenance.UpdateOrder(input); err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке обновления записи")
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) deleteOrder(c *gin.Context) {
	var input structures.OrderList
	c.Set("input", &input)

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, err, http.StatusBadRequest, "Неверный формат переданных данных")
		return
	}

	if err := h.services.OngoingMaintenance.DeleteOrder(input); err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при попытке удаления записи")
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) clear(c *gin.Context) {
	if err := h.services.OngoingMaintenance.Clear(); err != nil {
		newResponseError(c, err, http.StatusInternalServerError, "Внутренняя ошибка сервера при выполнении метода очистки")
		return
	}

	c.Status(http.StatusOK)
}
