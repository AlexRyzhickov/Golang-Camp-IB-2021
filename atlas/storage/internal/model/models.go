package models

type Note struct {
	Service string `gorm:"primaryKey"`
	Mode    bool
}
