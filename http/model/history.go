package model

import "time"

type Log struct {
	Id       string    `json:"id"`
	RunningT time.Time `json:"running_t"`
}
