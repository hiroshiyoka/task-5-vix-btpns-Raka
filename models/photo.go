package models

type Photo struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Title    string `gorm:"type:text" json:"caption"`
	Caption  string `gorm:"type:varchar(300)" json:"title"`
	PhotoUrl string `gorm:"string" json:"photourl"`
	UserID   string `json:"userid"`
}
