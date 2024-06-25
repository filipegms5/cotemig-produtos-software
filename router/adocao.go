package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupAdocaoRoutes(router *gin.Engine, adocaoRepository *repositories.AdocaoRepository) {
	router.GET("/adocoes", func(c *gin.Context) {
		controllers.GetAdocoes(c, adocaoRepository)
	})
	router.GET("/adocao/id/:id", func(c *gin.Context) {
		controllers.GetAdocao(c, adocaoRepository)
	})
	router.POST("/adocao", func(c *gin.Context) {
		controllers.CreateAdocao(c, adocaoRepository)
	})
	router.PUT("/adocao/id/:id", func(c *gin.Context) {
		controllers.UpdateAdocao(c, adocaoRepository)
	})
	router.DELETE("/adocao/id/:id", func(c *gin.Context) {
		controllers.DeleteAdocao(c, adocaoRepository)
	})
	router.GET("/adocao/adotante/:adotante", func(c *gin.Context) {
		controllers.GetAdocoesByAdotanteId(c, adocaoRepository)
	})
	router.GET("/adocao/idoso/:idoso", func(c *gin.Context) {
		controllers.GetAdocoesByIdosoId(c, adocaoRepository)
	})
}
