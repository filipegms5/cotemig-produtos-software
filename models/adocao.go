package models

type Adocao struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Data        string `gorm:"type:DATE ;NOT NULL" json:"data" binding:"required"`
	Status      string `gorm:"type:varchar(255);NOT NULL" json:"status" binding:"required"`
	Observacoes string `gorm:"type:varchar(255);" json:"observacoes"`
	AdotanteID  uint   `json:"adotanteID" gorm:"foreignKey:Adotante;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IdosoID     uint   `json:"idosoID" gorm:"foreignKey:Idoso;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
