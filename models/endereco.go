package models

type Endereco struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Logradouro  string `gorm:"type:varchar(255);NOT NULL" json:"logradouro" binding:"required"`
	Numero      string `gorm:"type:varchar(255);NOT NULL" json:"numero" binding:"required"`
	Complemento string `gorm:"type:varchar(255)" json:"complemento"`
	Cep         string `gorm:"type:varchar(255);NOT NULL" json:"cep" binding:"required"`
	Uf          string `gorm:"type:varchar(255);NOT NULL" json:"uf" binding:"required"`
	Cidade      string `gorm:"type:varchar(255);NOT NULL" json:"cidade" binding:"required"`
}
