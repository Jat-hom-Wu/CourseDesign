package router

import (
	"Global/controller"
	// "Global/cors"
	// "time"
	"github.com/gin-gonic/gin"
)
func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   w := c.Writer
	   // 处理js-ajax跨域问题
	   w.Header().Set("Access-Control-Allow-Origin", "http://159.75.2.47:8888") //允许访问所有域
	   w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
	   w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	   w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
	   w.Header().Set("Access-Control-Allow-Credentials", "true")
	//    w.Header().Add("X-Content-Type-Options", "nosniff")
	   c.Next()
	}
 }
func RountersInit() *gin.Engine {
	r := gin.Default()
	r.Use(AccessJsMiddleware())
	// r.Use(cors.Middleware(cors.Config{
	// 	Origins:        "http://159.75.2.47:8888",
	// 	Methods:        "GET, PUT, POST, DELETE",
	// 	RequestHeaders: "Origin, Authorization, Content-Type, Access-Token",
	// 	ExposedHeaders: "",
	// 	MaxAge: 50 * time.Second,
	// 	Credentials: true,
	// 	ValidateHeaders: false,
	// }))

	r.LoadHTMLGlob("templates/*")
	r.Static("/picture","./picture")
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
		group.GET("/home", controller.JwtMiddleWare) //show data page
		group.POST("/LoginCGISQL.cgi", controller.HandleLoginCGI)
		group.POST("/RegisterCGISQL.cgi", controller.HandleRegisterCGI)
	}

	return r
}
