package main

import (
	"flag"
	"ticket-system/config"
	controller "ticket-system/controllers"
	"ticket-system/databases/mysql"
	"ticket-system/databases/redis"
)

var configFile = flag.String("config", "app.yml", "config file")

func main() {
	flag.Parse()

	config.Init(*configFile)
	mysql.Init()
	redis.Init()
	controller.Init()
}
