package user

import "fmt"

// Money subdomain
type Money struct {
	Amount int64 `json:"amount"` //amount
}

func (m *Money) String() string {
	return fmt.Sprintf("%.2f", float64(m.Amount/100))
}
