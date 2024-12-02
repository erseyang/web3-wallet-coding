package repository

import (
	"com.wallet/coding/config"
	"com.wallet/coding/domain/user"
	"com.wallet/coding/infrastructure/postgres"
	"database/sql"
	"fmt"
)

// Check implements
var _ user.IUserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (i UserRepositoryImpl) UserList() ([]user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i UserRepositoryImpl) QueryUserByUserId(userId string) (*user.User, error) {
	query := "SELECT * FROM t_user WHERE user_id = ?"
	var u user.User
	err := i.db.QueryRow(query, userId).Scan(&u.UserId, &u.Name)
	if err != nil {
		return nil, err
	}
	return &u, nil
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

// Add
// user persistence
func (i UserRepositoryImpl) Add(user user.User) (int64, error) {
	fmt.Println("add had execute...")
	query := "INSERT INTO t_user (user_id, name, account, salt, password) VALUES (?, ?, ?, ?, ?)"
	affected, err := i.db.Exec(query, user.UserId, user.Name, user.Account, user.Salt, user.Password)
	if err != nil {
		return 0, err
	}
	var result int64
	result, err = affected.RowsAffected()
	return result, err
}

//func (i UserRepositoryImpl) UserExists(userId string) (bool, error) {
//	fmt.Println("user exist had execute")
//	query := "SELECT * FROM t_user WHERE user_id = ?"
//	rows, err := i.db.Query(query, userId)
//	if err != nil {
//		return false, err
//	}
//
//	return false, errors.New("err")
//}

func (i UserRepositoryImpl) CheckAmount() {

}
