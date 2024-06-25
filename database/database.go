package database

import (
	"github.com/filipegms5/cotemig-projeto-software/models"

	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Endereco{}, &models.Usuario{}, &models.Adotante{}, &models.Idoso{}, &models.Monitor{}, &models.Adocao{}, &models.Monitoria{})
}
