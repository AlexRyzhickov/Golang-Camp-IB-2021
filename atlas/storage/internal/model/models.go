package models

import "time"

type Note struct {
	Service string `gorm:"primaryKey"`
	Mode    bool
}

type Service struct {
	ServiceName          string
	ServiceDesc          string
	ServiceUptime        time.Time
	ServiceCountRequests uint
}

type Msg struct {
	Id      string
	Command string
}

type StorageResponse struct {
	Id    string
	Value interface{}
}
