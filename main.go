package main

import(
	"Global/router"
)

func main(){
	r := router.RountersInit()
	r.Run("127.0.0.1:9527")
}