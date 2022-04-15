package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

//show data page
func HandleHome(c *gin.Context){
	c.HTML(http.StatusOK, "home.html",nil)
}

//get data interface
func HandleData(c *gin.Context){
	c.JSON(http.StatusOK,"handle data api")
}

//inedex
func HandleIndex(c *gin.Context){
	c.HTML(http.StatusOK, "index.html",nil)
}

//login
func HandleLogin(c *gin.Context){
	c.HTML(http.StatusOK, "log.html",nil)
}
//login fail
func HandleLoginFail(c *gin.Context){
	c.HTML(http.StatusOK, "logError.html",nil)
}
//register
func HandleRegister(c *gin.Context){
	c.HTML(http.StatusOK, "register.html",nil)
}
//register fail
func HandleRegisterFail(c *gin.Context){
	c.HTML(http.StatusOK, "registerError.html",nil)
}

func HandleLoginCGI(c *gin.Context){
	user := c.PostForm("user")
	password := c.PostForm("password")
	re := user + password
	c.JSON(200,re)
}

func HandleRegisterCGI(c *gin.Context){
	user := c.PostForm("user")
	password := c.PostForm("password")
	re := user + password
	c.JSON(200,re)
}