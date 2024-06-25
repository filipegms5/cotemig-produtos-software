package models

type Usuario struct {
	Id             uint   `json:"id" gorm:"primaryKey"`
	Nome           string `gorm:"type:varchar(255);NOT NULL" json:"nome" binding:"required"`
	Email          string `gorm:"type:varchar(255);NOT NULL" json:"email" binding:"required"`
	Senha          string `gorm:"type:varchar(255);NOT NULL" json:"senha" binding:"required"`
	Telefone       string `gorm:"type:varchar(255);NOT NULL" json:"telefone" binding:"required"`
	Cpf            string `gorm:"type:varchar(255);NOT NULL" json:"cpf" binding:"required"`
	DataNascimento string `gorm:"type:DATE ;NOT NULL" json:"dataNascimento" binding:"required"`
	EnderecoID     uint   `gorm:"column:endereco_id;NOT NULL" json:"enderecoID" binding:"required"`
}
