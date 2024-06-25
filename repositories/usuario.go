package repositories

import (
	"github.com/filipegms5/cotemig-projeto-software/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{db}
}

func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
	return r.db.Create(usuario).Error
}
func (r *UsuarioRepository) Update(usuario *models.Usuario) error {
	return r.db.Save(usuario).Error
}

func (r *UsuarioRepository) Delete(usuario *models.Usuario) error {
	return r.db.Delete(usuario).Error
}

func (r *UsuarioRepository) FindById(id uint) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.First(&usuario, id).Error
	return &usuario, err
}

func (r *UsuarioRepository) FindAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Find(&usuarios).Error
	return usuarios, err
}

func (r *UsuarioRepository) FindByEmail(email string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("email = ?", email).First(&usuario).Error
	return &usuario, err
}

func (r *UsuarioRepository) FindByCpf(cpf string) (*models.Usuario, error) {
	var usuario models.Usuario
	err := r.db.Where("cpf = ?", cpf).First(&usuario).Error
	return &usuario, err
}

func (r *UsuarioRepository) FindByName(name string) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.db.Where("nome = ?", name).Find(&usuarios).Error
	return usuarios, err
}
