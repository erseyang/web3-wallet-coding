package transfer

import "errors"

type Transfer struct {
	FromUser   string `json:"from_user"`
	ToUser     string `json:"to_user"`
	TransferNo string `json:"transfer_no"` // transfer no

}

func (t *Transfer) Send(fromUserId, toUserId string, amount int) error {
	return errors.New("")
}
