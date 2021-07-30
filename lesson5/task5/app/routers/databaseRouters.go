package routers

import (
	"app/controllers"
	"app/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterDatabaseEndpoints(router *gin.RouterGroup, service usecase.Usecase) {

	handler := controllers.NewAppControllers(service)

	db := router.Group("/db")
	{
		db.POST("", handler.CreateNode)
		db.GET("", handler.GetNodes)
	}

}
