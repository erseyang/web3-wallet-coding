package repository

import (
	"com.wallet/coding/config"
	"com.wallet/coding/domain/user"
	"com.wallet/coding/infrastructure/postgres"
	"database/sql"
	"errors"
	"fmt"
)

// Check implements
var _ user.IUserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: postgres.Instance(
			config.Instance.PostgreSql.DbHost,
			config.Instance.PostgreSql.DbUser,
			config.Instance.PostgreSql.DbPass,
			config.Instance.PostgreSql.DbName,
			config.Instance.PostgreSql.DbPort,
		),
	}
}

func (i UserRepositoryImpl) Add(user user.User) (int64, error) {
	fmt.Println("add had execute...")
	return 0, errors.New("err")
}

func (i UserRepositoryImpl) UserExists(userId string) (bool, error) {
	fmt.Println("user exist had execute")
	return false, errors.New("err")
}

func (i UserRepositoryImpl) CheckAmount() {

}
