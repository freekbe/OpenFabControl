package model

import (
	"time"
)

type Machine_controller struct {
	ID			int			`json:"id"`
	UUID 		string		`json:"uuid"`
	Approved 	bool		`json:"approved"`
	CreatedAt	time.Time	`json:"created_at"`
}
