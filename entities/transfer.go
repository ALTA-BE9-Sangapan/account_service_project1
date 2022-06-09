package entities

import (
	"time"
)

type Transfer struct {
	ID              int
	SenderPhone     string
	ReceiverPhone   string
	TransferBalance int
	CreatedAt       time.Time
}
