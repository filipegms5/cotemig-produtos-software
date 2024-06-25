package router

import (
	"github.com/filipegms5/cotemig-projeto-software/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	adocaoRepository := repositories.NewAdocaoRepository(db)
	adotanteRepository := repositories.NewAdotanteRepository(db)
	enderecoRepository := repositories.NewEnderecoRepository(db)
	idosoRepository := repositories.NewIdosoRepository(db)
	monitorRepository := repositories.NewMonitorRepository(db)
	monitoriaRepository := repositories.NewMonitoriaRepository(db)
	usuarioRepository := repositories.NewUsuarioRepository(db)

	router := gin.Default()

	SetupUsuarioRoutes(router, usuarioRepository)

	SetupAdocaoRoutes(router, adocaoRepository)

	SetupAdotanteRoutes(router, adotanteRepository)

	SetupEnderecoRoutes(router, enderecoRepository)

	SetupIdosoRoutes(router, idosoRepository)

	SetupMonitorRoutes(router, monitorRepository)

	SetupMonitoriaRoutes(router, monitoriaRepository)

	return router
}
