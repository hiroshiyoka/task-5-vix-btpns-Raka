package models

type User struct {
	ID       int64   `gorm:"primaryKey" json:"id"`
	Username string  `gorm:"varchar(300)" json:"username"`
	Email    string  `gorm:"varchar(300)" json:"email"`
	Password string  `gorm:"varchar(300)" json:"password"`
	Photos   []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
}
