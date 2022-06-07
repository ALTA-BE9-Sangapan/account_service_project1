package entities

import (
	"time"
)

type Users struct {
	user_phone string
	user_id    string
	saldo      int
	password   string
	date_birth time.Time
	gender     string
	alamat     string
}
