package models

type Monitoria struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Data        string `gorm:"type:DATE ;NOT NULL" json:"data" binding:"required"`
	Observacoes string `gorm:"type:varchar(255);" json:"observacoes"`
	MonitorID   uint   `json:"monitor_id" gorm:"foreignKey:MonitorID"`
	AdocaoID    uint   `json:"adocao_id" gorm:"foreignKey:AdocaoID"`
}
