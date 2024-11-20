package dto

type UserDto struct {
}

type UserLoginDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
