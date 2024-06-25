package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupAdotanteRoutes(router *gin.Engine, adotanteRepository *repositories.AdotanteRepository) {

	router.GET("/adotantes", func(c *gin.Context) {
		controllers.GetAdotantes(c, adotanteRepository)
	})
	router.GET("/adotante/id/:id", func(c *gin.Context) {
		controllers.GetAdotante(c, adotanteRepository)
	})
	router.POST("/adotante", func(c *gin.Context) {
		controllers.CreateAdotante(c, adotanteRepository)
	})
	router.PUT("/adotante/id/:id", func(c *gin.Context) {
		controllers.UpdateAdotante(c, adotanteRepository)
	})
	router.DELETE("/adotante/id/:id", func(c *gin.Context) {
		controllers.DeleteAdotante(c, adotanteRepository)
	})
	router.GET("/adotante/usuarioid/:usuarioid", func(c *gin.Context) {
		controllers.GetAdotanteByUsuarioID(c, adotanteRepository)
	})
}
