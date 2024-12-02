package user

type IUserRepository interface {
	CheckAmount()
	Add(user User) (int64, error)
	QueryUserByUserId(userId string) (*User, error)
	UserList() ([]User, error)
}
