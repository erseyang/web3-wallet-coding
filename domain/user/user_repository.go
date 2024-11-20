package user

type IUserRepository interface {
	UserExists(userId string) (bool, error)
	CheckAmount()
	Add(user User) (int64, error)
}
