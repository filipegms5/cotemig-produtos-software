package models

type Adotante struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Resumo    string `gorm:"type:varchar(255);NOT NULL" json:"resumo" binding:"required"`
	UsuarioID uint   `json:"usuario_id" gorm:"foreignKey:Usuario"`
}
