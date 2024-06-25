package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type IdosoRepository struct {
	db *gorm.DB
}

func NewIdosoRepository(db *gorm.DB) *IdosoRepository {
	return &IdosoRepository{db}
}

func (r *IdosoRepository) Create(idoso *models.Idoso) error {
	return r.db.Create(idoso).Error
}

func (r *IdosoRepository) Update(idoso *models.Idoso) error {
	return r.db.Save(idoso).Error
}

func (r *IdosoRepository) Delete(idoso *models.Idoso) error {
	return r.db.Delete(idoso).Error
}

func (r *IdosoRepository) FindById(id uint) (*models.Idoso, error) {
	var idoso models.Idoso
	err := r.db.First(&idoso, id).Error
	return &idoso, err
}

func (r *IdosoRepository) FindAll() ([]models.Idoso, error) {
	var idosos []models.Idoso
	err := r.db.Find(&idosos).Error
	return idosos, err
}

func (r *IdosoRepository) FindByUsuarioId(id uint) (*models.Idoso, error) {
	var idoso models.Idoso
	err := r.db.Where("usuario_id = ?", id).First(&idoso).Error
	return &idoso, err
}
