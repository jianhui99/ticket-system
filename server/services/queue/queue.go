package queue

import (
	"ticket-system/databases/mysql"
	"ticket-system/models/queue"
)

func Create(u *queue.Queue) error {
	result := mysql.MysqlDB.Create(u)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
