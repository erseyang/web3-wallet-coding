package transfer

import (
	"errors"
	"time"
)

type Transfer struct {
	FromUser   string    `json:"from_user"`
	ToUser     string    `json:"to_user"`
	TransferNo string    `json:"transfer_no"` // transfer no
	Money      Money     `json:"money"`
	CreateTs   time.Time `json:"create_ts"`
}

func (t *Transfer) Send(fromUserId, toUserId string, amount int) error {
	return errors.New("")
}
