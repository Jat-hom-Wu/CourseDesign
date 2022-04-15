package main

import(
	"net"
	"log"
	"Global/terminal_server"
	"Global/dao"
)

func main(){
	log.Println("start")
	dao.MySQLInit()
	listener,err := net.Listen("tcp",":8888")
	if err != nil{
		log.Println("terminal server listen failed:",err)
	}
	for{
		conn,_ := listener.Accept()
		go terminal_server.Service(conn)
	}
}