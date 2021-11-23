package models

import "time"

type Service struct {
	ServiceName          string
	ServiceDesc          string
	ServiceUptime        time.Time
	ServiceCountRequests uint
}

type Msg struct {
	Id      string
	Command string
	Value   interface{}
}
