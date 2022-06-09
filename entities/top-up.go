package entities

import "time"

type TopUp struct {
	ID           int
	TopUpPhone   string
	TopUpBalance int
	CreatedAt    time.Time
}
