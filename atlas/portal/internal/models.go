package models

import "time"

type Service struct {
	ServiceName          string
	ServiceDesc          string
	ServiceUptime        time.Time
	ServiceCountRequests uint
}
