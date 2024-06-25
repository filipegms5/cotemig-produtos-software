package router

import (
	"github.com/filipegms5/cotemig-projeto-software/controllers"
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
)

func SetupUsuarioRoutes(router *gin.Engine, usuarioRepository *repositories.UsuarioRepository) {
	router.GET("/usuarios", func(c *gin.Context) {
		controllers.GetUsuarios(c, usuarioRepository)
	})
	router.GET("/usuario/id/:id", func(c *gin.Context) {
		controllers.GetUsuario(c, usuarioRepository)
	})
	router.POST("/usuario", func(c *gin.Context) {
		controllers.CreateUsuario(c, usuarioRepository)
	})
	router.PUT("/usuario/id/:id", func(c *gin.Context) {
		controllers.UpdateUsuario(c, usuarioRepository)
	})
	router.DELETE("/usuario/id/:id", func(c *gin.Context) {
		controllers.DeleteUsuario(c, usuarioRepository)
	})
	router.GET("/usuario/cpf/:cpf", func(c *gin.Context) {
		controllers.GetUsuarioByCpf(c, usuarioRepository)
	})
	router.GET("/usuario/email/:email", func(c *gin.Context) {
		controllers.GetUsuarioByEmail(c, usuarioRepository)
	})
	router.GET("/usuario/name/:name", func(c *gin.Context) {
		controllers.GetUsuarioByName(c, usuarioRepository)
	})
}
