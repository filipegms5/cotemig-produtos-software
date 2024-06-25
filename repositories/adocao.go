package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type AdocaoRepository struct {
	db *gorm.DB
}

func NewAdocaoRepository(db *gorm.DB) *AdocaoRepository {
	return &AdocaoRepository{db}
}

func (r *AdocaoRepository) Create(adocao *models.Adocao) error {
	return r.db.Create(adocao).Error
}

func (r *AdocaoRepository) Update(adocao *models.Adocao) error {
	return r.db.Save(adocao).Error
}

func (r *AdocaoRepository) Delete(adocao *models.Adocao) error {
	return r.db.Delete(adocao).Error
}

func (r *AdocaoRepository) FindById(id uint) (*models.Adocao, error) {
	var adocao models.Adocao
	err := r.db.First(&adocao, id).Error
	return &adocao, err
}

func (r *AdocaoRepository) FindAll() ([]models.Adocao, error) {
	var adocoes []models.Adocao
	err := r.db.Find(&adocoes).Error
	return adocoes, err
}

func (r *AdocaoRepository) FindByAdotanteId(id uint) ([]models.Adocao, error) {
	var adocoes []models.Adocao
	err := r.db.Where("adotante_id = ?", id).Find(&adocoes).Error
	return adocoes, err
}

func (r *AdocaoRepository) FindByIdosoId(id uint) ([]models.Adocao, error) {
	var adocoes []models.Adocao
	err := r.db.Where("idoso_id = ?", id).Find(&adocoes).Error
	return adocoes, err
}
