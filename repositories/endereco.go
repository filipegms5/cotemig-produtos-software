package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type EnderecoRepository struct {
	db *gorm.DB
}

func NewEnderecoRepository(db *gorm.DB) *EnderecoRepository {
	return &EnderecoRepository{db}
}

func (r *EnderecoRepository) Create(endereco *models.Endereco) error {
	return r.db.Create(endereco).Error
}

func (r *EnderecoRepository) FindById(id uint) (*models.Endereco, error) {
	var endereco models.Endereco
	err := r.db.First(&endereco, id).Error
	return &endereco, err
}

func (r *EnderecoRepository) FindAll() ([]models.Endereco, error) {
	var enderecos []models.Endereco
	err := r.db.Find(&enderecos).Error
	return enderecos, err
}

func (r *EnderecoRepository) Update(endereco *models.Endereco) error {
	return r.db.Save(endereco).Error
}

func (r *EnderecoRepository) Delete(endereco *models.Endereco) error {
	return r.db.Delete(endereco).Error
}

func (r *EnderecoRepository) FindByCep(cep string) (*models.Endereco, error) {
	var endereco models.Endereco
	err := r.db.Where("cep = ?", cep).First(&endereco).Error
	return &endereco, err
}
