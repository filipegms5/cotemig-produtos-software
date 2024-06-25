package models

type Idoso struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Observacoes string `gorm:"type:varchar(255);" json:"observacoes"`
	UsuarioID   uint   `json:"usuarioID" gorm:"foreignKey:Usuario;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
