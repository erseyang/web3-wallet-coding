package user

import "com.wallet/coding/infrastructure/utils"

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Wallet   Wallet `json:"wallet"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func (u *User) TableName() string {
	return "t_user"
}

func (u *User) CheckPasswordMatch(password string) bool {
	temp := &User{
		Password: password,
		Salt:     u.Salt,
	}
	if u.Password == temp.Sha256Password() {
		return true
	}
	return false
}

func (u *User) Sha256Password() string {
	return utils.Sha256(u.Password + u.Salt)
}
