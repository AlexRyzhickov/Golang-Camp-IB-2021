package models

import (
	"sync"
	"time"
)

type Service struct {
	sync.RWMutex
	ServiceDesc          string
	ServiceUptime        time.Time
	ServiceCountRequests uint
}

type Msg struct {
	Id      string
	Command string
	Value   interface{}
}
