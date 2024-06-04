package user

import (
	"ticket-system/databases/mysql"
	"ticket-system/models/user"
)

func Create(u *user.User) error {
	result := mysql.MysqlDB.Create(u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
