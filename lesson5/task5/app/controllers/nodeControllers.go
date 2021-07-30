package controllers

import (
	"app/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ctrl *AppControllers) GetModemInformations(ctx *gin.Context) {

	var data interface{}
	var proxy models.Nodes
	var hostname string = ctx.Param("hostname")
	var path string = "/v2/modem/info"

	if err := ctrl.uc.FindProxyByHostname(&proxy, hostname); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := ctrl.uc.Get(proxy.Address, path, &data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (ctrl *AppControllers) GetInternetProvider(ctx *gin.Context) {

	var data interface{}
	var proxy models.Nodes
	var hostname string = ctx.Param("hostname")
	var path string = "/v2/modem/provider"

	if err := ctrl.uc.FindProxyByHostname(&proxy, hostname); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := ctrl.uc.Get(proxy.Address, path, &data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (ctrl *AppControllers) GetMonitoringStatus(ctx *gin.Context) {

	var data interface{}
	var proxy models.Nodes
	var hostname string = ctx.Param("hostname")
	var path string = "/v2/modem/status"

	if err := ctrl.uc.FindProxyByHostname(&proxy, hostname); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := ctrl.uc.Get(proxy.Address, path, &data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (ctrl *AppControllers) RebootModem(ctx *gin.Context) {

	var proxy models.Nodes
	var hostname string = ctx.Param("hostname")
	var path string = "/v2/modem/reboot"

	if err := ctrl.uc.FindProxyByHostname(&proxy, hostname); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := ctrl.uc.RebootModem(proxy.Address, path); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Modem reboot started"})
}
