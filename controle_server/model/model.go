package model

import (
	"time"
)

type Machine_controller struct {
	ID						int			`json:"id"`
	UUID 					string		`json:"uuid"`
	TYPE					string		`json:"type"`
	ZONE 					string		`json:"zone"`
	NAME 					string		`json:"name"`
	MANUAL					string		`json:"manual"`
	PRICE_BOOKING_IN_EUR	float64		`json:"price_booking_in_eur"`
	PRICE_USAGE_IN_EUR		float64		`json:"price_usage_in_eur"`
	Approved 				bool		`json:"approved"`
	CreatedAt				time.Time	`json:"created_at"`
}
