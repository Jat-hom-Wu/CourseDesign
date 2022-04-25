package router

import (
	"Global/controller"

	"github.com/gin-gonic/gin"
)

func RountersInit() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/picture","./picture")
	// r.GET("/pid", HandlePid)	//
	r.GET("/data", controller.HandleData) //get all data

	group := r.Group("/course")
	//语句块
	{
		//起始页
		group.GET("/index", controller.HandleIndex)
		//登录
		group.GET("/login", controller.HandleLogin)
		//注册
		group.GET("/register", controller.HandleRegister)
		//登录检测
		group.GET("/loginfail", controller.HandleLoginFail)
		//注册检测
		group.GET("/registerfail", controller.HandleRegisterFail)
		//展示页面
		group.GET("/home", controller.JwtMiddleWare, controller.HandleHome) //show data page
		group.POST("/LoginCGISQL.cgi", controller.HandleLoginCGI)
		group.POST("/RegisterCGISQL.cgi", controller.HandleRegisterCGI)
	}

	return r
}
