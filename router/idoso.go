package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupIdosoRoutes(router *gin.Engine, idosoRepository *repositories.IdosoRepository) {
	router.GET("/idosos", func(c *gin.Context) {
		controllers.GetIdosos(c, idosoRepository)
	})
	router.GET("/idoso/id/:id", func(c *gin.Context) {
		controllers.GetIdoso(c, idosoRepository)
	})
	router.POST("/idoso", func(c *gin.Context) {
		controllers.CreateIdoso(c, idosoRepository)
	})
	router.PUT("/idoso/id/:id", func(c *gin.Context) {
		controllers.UpdateIdoso(c, idosoRepository)
	})
	router.DELETE("/idoso/id/:id", func(c *gin.Context) {
		controllers.DeleteIdoso(c, idosoRepository)
	})
	router.GET("/idoso/usuarioid/:usuarioid", func(c *gin.Context) {
		controllers.GetIdosoByUsuarioId(c, idosoRepository)
	})
}
