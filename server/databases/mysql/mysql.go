package mysql

import (
	"fmt"
	"strconv"
	"ticket-system/config"
	"ticket-system/models/order"
	"ticket-system/models/queue"
	"ticket-system/models/ticket"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Database.MySqlConf.Username, config.Conf.Database.MySqlConf.Password, config.Conf.Database.MySqlConf.Host, strconv.Itoa(config.Conf.Database.MySqlConf.Port), config.Conf.Database.MySqlConf.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&ticket.Ticket{}, &order.Order{}, &queue.Queue{})

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 全局变量
	MysqlDB = db

	fmt.Println("mysql init success")
}

func GetDB() *gorm.DB {
	return MysqlDB
}

func Close() {
	sqlDB, err := MysqlDB.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.Close()
	fmt.Println("mysql close success")
}
