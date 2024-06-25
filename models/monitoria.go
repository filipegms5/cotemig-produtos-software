package models

type Monitoria struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Data        string `gorm:"type:DATE ;NOT NULL" json:"data" binding:"required"`
	Observacoes string `gorm:"type:varchar(255);" json:"observacoes"`
	MonitorID   uint   `json:"monitorID" gorm:"foreignKey:MonitorID"`
	AdocaoID    uint   `json:"adocaoID" gorm:"foreignKey:AdocaoID"`
}
