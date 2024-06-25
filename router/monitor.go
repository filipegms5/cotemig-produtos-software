package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupMonitorRoutes(router *gin.Engine, monitorRepository *repositories.MonitorRepository) {
	router.GET("/monitores", func(c *gin.Context) {
		controllers.GetMonitores(c, monitorRepository)
	})
	router.GET("/monitor/id/:id", func(c *gin.Context) {
		controllers.GetMonitor(c, monitorRepository)
	})
	router.POST("/monitor", func(c *gin.Context) {
		controllers.CreateMonitor(c, monitorRepository)
	})
	router.PUT("/monitor/id/:id", func(c *gin.Context) {
		controllers.UpdateMonitor(c, monitorRepository)
	})
	router.DELETE("/monitor/id/:id", func(c *gin.Context) {
		controllers.DeleteMonitor(c, monitorRepository)
	})
	router.GET("/monitor/usuarioid/:usuarioid", func(c *gin.Context) {
		controllers.GetMonitorByUsuarioId(c, monitorRepository)
	})
}
