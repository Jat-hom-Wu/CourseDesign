package main

import(
	"Global/router"
	"Global/dao"
)

func main(){
	dao.MySQLInit()
	dao.RedisInit()
	r := router.RountersInit()
	r.Run(":9527")
}