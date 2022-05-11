package dao

//初始化数据库
//redis设置maxmemory为20M,轮训策略为lru

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
	"log"
	"Global/models"
)


func MySQLInit() error{
	dsn := "root:12345@tcp(127.0.0.1:3306)/data_monitor?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Println("mysql init:open database failed:",err)
		return err
	}
	models.DB = db
	return nil
}

func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	models.RDB = rdb
}