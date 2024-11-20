package user

type Wallet struct {
	Money   Money  `json:"money"`
	Account string `json:"account"`
}
