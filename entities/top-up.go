package entities

import "time"

type Top_up struct {
	top_up_id  int
	user_id    int
	saldo      int
	status     string
	created_at time.Time
}
