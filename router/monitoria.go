package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupMonitoriaRoutes(router *gin.Engine, monitoriaRepository *repositories.MonitoriaRepository) {
	router.GET("/monitorias", func(c *gin.Context) {
		controllers.GetMonitorias(c, monitoriaRepository)
	})
	router.GET("/monitoria/id/:id", func(c *gin.Context) {
		controllers.GetMonitoria(c, monitoriaRepository)
	})
	router.POST("/monitoria", func(c *gin.Context) {
		controllers.CreateMonitoria(c, monitoriaRepository)
	})
	router.PUT("/monitoria/id/:id", func(c *gin.Context) {
		controllers.UpdateMonitoria(c, monitoriaRepository)
	})
	router.DELETE("/monitoria/id/:id", func(c *gin.Context) {
		controllers.DeleteMonitoria(c, monitoriaRepository)
	})
	router.GET("/monitoria/monitorid/:monitorid", func(c *gin.Context) {
		controllers.GetMonitoriasByMonitorId(c, monitoriaRepository)
	})

}
