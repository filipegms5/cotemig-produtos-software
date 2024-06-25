package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupEnderecoRoutes(router *gin.Engine, enderecoRepository *repositories.EnderecoRepository) {
	router.GET("/enderecos", func(c *gin.Context) {
		controllers.GetEnderecos(c, enderecoRepository)
	})
	router.GET("/endereco/id/:id", func(c *gin.Context) {
		controllers.GetEndereco(c, enderecoRepository)
	})
	router.POST("/endereco", func(c *gin.Context) {
		controllers.CreateEndereco(c, enderecoRepository)
	})
	router.PUT("/endereco/id/:id", func(c *gin.Context) {
		controllers.UpdateEndereco(c, enderecoRepository)
	})
	router.DELETE("/endereco/id/:id", func(c *gin.Context) {
		controllers.DeleteEndereco(c, enderecoRepository)
	})
	router.GET("/endereco/cep/:cep", func(c *gin.Context) {
		controllers.GetEnderecoByCep(c, enderecoRepository)
	})

}
