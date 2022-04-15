package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"Global/models"
	"time"
	"Global/jwt_service"
	"log"
	"github.com/dgrijalva/jwt-go"
)

//show data page
func HandleHome(c *gin.Context){
	_,err := c.Cookie("token")
	if err != nil{
		log.Println("not receive token")
		c.Redirect(http.StatusFound, "/course/login")
		return
	}
	c.HTML(http.StatusOK,"home.html",nil)
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
			//token
			nowTime := time.Now()
			expireTime := nowTime.Add(600 * time.Second)	//token的过期时间，header中以设置过期时间，因此此处没意义
			issuer := "frank"
			cla := jwt_service.Claims{
				Password: password,
				Username: user,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expireTime.Unix(),
					Issuer:    issuer,
				},
			}
			token,err := jwt_service.GenerateToken(cla)
			if err != nil{
				log.Println("generate token falied:",err)
			}
			c.SetCookie("token", token, 600, "/", "159.75.2.47", false, false)
			c.Redirect(http.StatusFound, "/course/home")
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