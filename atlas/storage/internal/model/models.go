package models

import (
	"sync"
	"time"
)

type Note struct {
	Service string `gorm:"primaryKey"`
	Mode    bool
}

type Service struct {
	sync.RWMutex
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
