package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type MonitoriaRepository struct {
	db *gorm.DB
}

func NewMonitoriaRepository(db *gorm.DB) *MonitoriaRepository {
	return &MonitoriaRepository{db}
}

func (r *MonitoriaRepository) Create(monitoria *models.Monitoria) error {
	return r.db.Create(monitoria).Error
}

func (r *MonitoriaRepository) Update(monitoria *models.Monitoria) error {
	return r.db.Save(monitoria).Error
}

func (r *MonitoriaRepository) Delete(monitoria *models.Monitoria) error {
	return r.db.Delete(monitoria).Error
}

func (r *MonitoriaRepository) FindById(id uint) (*models.Monitoria, error) {
	var monitoria models.Monitoria
	err := r.db.First(&monitoria, id).Error
	return &monitoria, err
}

func (r *MonitoriaRepository) FindAll() ([]models.Monitoria, error) {
	var monitorias []models.Monitoria
	err := r.db.Find(&monitorias).Error
	return monitorias, err
}

func (r *MonitoriaRepository) FindByMonitorId(id uint) ([]models.Monitoria, error) {
	var monitorias []models.Monitoria
	err := r.db.Where("monitor_id = ?", id).Find(&monitorias).Error
	return monitorias, err
}

func (r *MonitoriaRepository) FindByAdocaoId(id uint) (*models.Monitoria, error) {
	var monitoria models.Monitoria
	err := r.db.Where("adocao_id = ?", id).First(&monitoria).Error
	return &monitoria, err
}
