package controllers

import (
	"app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *AppControllers) CreateNode(ctx *gin.Context) {

	var nodes models.Nodes

	if err := ctx.BindJSON(&nodes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.uc.CreateProxy(&nodes); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

func (ctrl *AppControllers) GetNodes(ctx *gin.Context) {

	var nodes []models.Nodes

	if err := ctrl.uc.FindNodes(&nodes); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, nodes)

}
