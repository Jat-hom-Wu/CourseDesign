package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"Global/models"
)

//show data page
func HandleHome(c *gin.Context){
	c.HTML(http.StatusOK, "home.html",nil)
}

//get data interface
func HandleData(c *gin.Context){
	result,_ := models.HumidifierGetData()
	c.JSON(200, *result)
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
	result,err := models.UserFindData(user)
	if err != nil{
		c.JSON(505,"server error")
		return
	}else{
		if user == result.UserName && password == result.Password{
			c.HTML(http.StatusOK, "home.html", nil)
		}else{
			c.HTML(http.StatusOK, "logError.html", nil)
		}
	}
	
}

func HandleRegisterCGI(c *gin.Context){
	user := c.PostForm("user")
	password := c.PostForm("password")
	result,err := models.UserFindData(user)
	if err != nil{
		c.JSON(505,"server error")
		return
	}else{
		if result.UserName == ""{
			models.UserCreateData(user,password)
			c.HTML(http.StatusOK, "log.html", nil)
		}else{
			c.HTML(http.StatusOK, "registerError.html", nil)
		}
	}
}