package models

type Note struct {
	Service string `gorm:"primaryKey"`
	Mode    bool
}

type Msg struct {
	Id      string
	Command string
}

type StorageResponse struct {
	Id    string
	Value interface{}
}
