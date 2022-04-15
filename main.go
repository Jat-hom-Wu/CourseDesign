package main

import(
	"Global/router"
	"Global/dao"
)

func main(){
	dao.MySQLInit()
	r := router.RountersInit()
	r.Run(":9527")
}