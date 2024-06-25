package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type MonitorRepository struct {
	db *gorm.DB
}

func NewMonitorRepository(db *gorm.DB) *MonitorRepository {
	return &MonitorRepository{db}
}

func (r *MonitorRepository) Create(monitor *models.Monitor) error {
	return r.db.Create(monitor).Error
}

func (r *MonitorRepository) Update(monitor *models.Monitor) error {
	return r.db.Save(monitor).Error
}

func (r *MonitorRepository) Delete(monitor *models.Monitor) error {
	return r.db.Delete(monitor).Error
}

func (r *MonitorRepository) FindById(id uint) (*models.Monitor, error) {
	var monitor models.Monitor
	err := r.db.First(&monitor, id).Error
	return &monitor, err
}

func (r *MonitorRepository) FindAll() ([]models.Monitor, error) {
	var monitors []models.Monitor
	err := r.db.Find(&monitors).Error
	return monitors, err
}

func (r *MonitorRepository) FindByUsuarioId(id uint) (*models.Monitor, error) {
	var monitor models.Monitor
	err := r.db.Where("usuario_id = ?", id).First(&monitor).Error
	return &monitor, err
}
