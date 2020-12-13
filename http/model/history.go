package model

import "time"

type History struct {
	Id       string    `json:"id"`
	RunningT time.Time `json:"running_t"`
}
