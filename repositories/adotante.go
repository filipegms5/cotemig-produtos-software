package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type AdotanteRepository struct {
	db *gorm.DB
}

func NewAdotanteRepository(db *gorm.DB) *AdotanteRepository {
	return &AdotanteRepository{db}
}

func (r *AdotanteRepository) Create(adotante *models.Adotante) error {
	return r.db.Create(adotante).Error
}

func (r *AdotanteRepository) Update(adotante *models.Adotante) error {
	return r.db.Save(adotante).Error
}

func (r *AdotanteRepository) Delete(adotante *models.Adotante) error {
	return r.db.Delete(adotante).Error
}

func (r *AdotanteRepository) FindById(id uint) (*models.Adotante, error) {
	var adotante models.Adotante
	err := r.db.First(&adotante, id).Error
	return &adotante, err
}

func (r *AdotanteRepository) FindAll() ([]models.Adotante, error) {
	var adotantes []models.Adotante
	err := r.db.Find(&adotantes).Error
	return adotantes, err
}

func (r *AdotanteRepository) FindByUsuarioID(id uint) (*models.Adotante, error) {
	var adotante models.Adotante
	err := r.db.Where("usuario_id = ?", id).First(&adotante).Error
	return &adotante, err
}
