package routers

import (
	"app/controllers"
	"app/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterNodeEndpoints(router *gin.RouterGroup, service usecase.Usecase) {

	handler := controllers.NewAppControllers(service)

	node := router.Group("/nodes")
	{
		node.GET("/:hostname/modem/info", handler.GetModemInformations)
		node.GET("/:hostname/modem/net/provider", handler.GetInternetProvider)
		node.GET("/:hostname/modem/monitoring/status", handler.GetMonitoringStatus)
		node.POST("/:hostname/modem/reset", handler.RebootModem)
	}

}
