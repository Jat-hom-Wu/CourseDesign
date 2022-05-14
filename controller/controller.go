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

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
//Code: 1 2 3 4 5(toekn failed) 6(token success) 7(get cookie failed)

type Request struct{
	User string `json:user`
	Password string `json:password`
}


//show data page
func HandleHome(c *gin.Context){
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
	c.JSON(200,"login page")
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
	//这里做重定向的话就慢了，应该后端返回json，让前端来完成渲染html的。
	req := Request{}
	errBind := c.BindJSON(&req)
	if errBind != nil{
		log.Println("server login bind json failed:",errBind)
		c.JSON(200,Response{
			Code:3,
			Msg:"paramter error",
			Data:"",
		})
		return
	}
	result,err := models.UserFindData(req.User)
	log.Println("user:",req.User,"; password:",req.Password)
	if err != nil{
		c.JSON(505,"server error")
		return
	}else{
		if req.User != "" && req.User == result.UserName && req.Password == result.Password{
			//token
			nowTime := time.Now()
			expireTime := nowTime.Add(5 * time.Second)	//token的过期时间，需要区别于set_cookie中的过期时间
			issuer := "frank"
			cla := jwt_service.Claims{
				//token中最好不要放敏感信息
				Password: req.Password,
				Username: req.User,
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
			//refreshtoken generate
			refreshTokenExpireTime := nowTime.Add(30 * time.Second)	
			cla.StandardClaims.ExpiresAt = refreshTokenExpireTime.Unix()
			refreshToken,err := jwt_service.GenerateToken(cla)
			if err != nil{
				log.Println("generate token falied:",err)
			}
			c.SetCookie("refreshToken", refreshToken, 600, "/", "159.75.2.47", false, false)
			c.JSON(200, Response{
				Code:1,
				Msg:"",
				Data:"",
			})
		}else{
			c.JSON(200,Response{
				Code:2,
				Msg:"user or password error",
				Data:"",
			})
		}
	}
	
}

func HandleRegisterCGI(c *gin.Context){
	req := Request{}
	errBind := c.BindJSON(&req)
	if errBind != nil{
		log.Println("server register bind json failed:",errBind)
		c.JSON(200,Response{
			Code:3,
			Msg:"server bind json error",
			Data:"",
		})
		return
	}
	result,err := models.UserFindData(req.User)
	if err != nil{
		c.JSON(505,"server error")
		return
	}else{
		if result.UserName == ""{
			models.UserCreateData(req.User,req.Password)
			c.JSON(200, Response{
				Code:3,
				Msg:"",
				Data:"",
			})
		}else{
			c.JSON(200, Response{
				Code:4,
				Msg:"user name exist",
				Data:"",
			})
		}
	}
}

//Todo:重新颁发新token,用双token
func JwtMiddleWare(c *gin.Context){
	token,err := c.Cookie("token")
	refreshToken,refreshTokenErr := c.Cookie("refreshToken")
	if err != nil || refreshTokenErr != nil{
		log.Println("get cookie failed:",err,"; ",refreshTokenErr)
		c.JSON(200,Response{
			Code:7,
		})
		c.Abort()
		return
	}else{
		_,err := jwt_service.ParseToken(token)
		if err != nil{
			log.Println("token parse failed")
			claims,err := jwt_service.ParseToken(refreshToken)
			if err != nil{
				log.Println("refresh token failed too")
				c.JSON(200,Response{
					Code:5,
				})
				c.Abort()
				return
			}else{
				claToken := newClaims(5, claims.Username, claims.Password)
				claRefreshToken := newClaims(30, claims.Username, claims.Password)
				token,err := jwt_service.GenerateToken(claToken)
				if err != nil{
					log.Println("generate token falied:",err)
				}
				refreshToken,err := jwt_service.GenerateToken(claRefreshToken)
				if err != nil{
					log.Println("generate token falied:",err)
				}
				c.SetCookie("token", token, 600, "/", "159.75.2.47", false, false)
				c.SetCookie("refreshToken", refreshToken, 600, "/", "159.75.2.47", false, false)
				log.Println("refresh token parse success")
				c.JSON(200,Response{
					Code:6,
				})
			}
		}else{
			log.Println("token parse success")
			c.JSON(200,Response{
				Code:6,
			})
		}
	}
}

func newClaims(t int, user,password string) jwt_service.Claims{
	nowTime := time.Now()
	//time.Duration类型 不能直接和 int类型相乘，需要先将变量转换为time.Duration
	expireTime := nowTime.Add(time.Duration(t) * time.Second)	//token的过期时间，需要区别于set_cookie中的过期时间
	issuer := "frank"
	cla := jwt_service.Claims{
		//token中最好不要放敏感信息
		Password: password,
		Username: user,
		StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    issuer,
		},
	}
	return cla
}