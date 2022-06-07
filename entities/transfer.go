package entities

import (
	"time"
)

type Transfer struct {
	transfer_id         int
	user_phone          int
	user_phone_receiver int
	saldo               int
	status              string
	password            string
	created_at          time.Time
}
