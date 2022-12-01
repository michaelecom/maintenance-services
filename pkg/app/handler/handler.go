package handler

import (
	"github.com/gin-gonic/gin"
	"rimeks.ru/services/pkg/app/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	api := router.Group("api")
	{
		ongoing := api.Group("ongoing-maintenance")
		{
			ongoing.POST("/", h.createOrder)
			ongoing.GET("/", h.getAllOrders)
			ongoing.GET("/:service_market_id", h.getAllOrdersByServiceMarketID)
			ongoing.PUT("/", h.updateOrder)
			ongoing.DELETE("/", h.deleteOrder)

			ongoing.GET("/clear", h.clearOrders)

			logging := ongoing.Group("log")
			{
				logging.POST("/", h.createLog)
				logging.GET("/", h.getAllLogs)
				logging.GET("/:service_market_id", h.getAllLogsByServiceMarketID)

				logging.GET("/clear", h.clearLogs)
			}
		}
	}

	return router
}
