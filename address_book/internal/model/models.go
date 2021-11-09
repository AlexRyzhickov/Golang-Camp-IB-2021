package models

type Contact struct {
	Phone   string `gorm:"primaryKey"`
	Name    string
	Address string
}
