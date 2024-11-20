package dto

type TransferDto struct {
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	Amount     int64  `json:"amount"`
}

type TransferDepositDto struct {
	Amount int64 `json:"amount"`
}
