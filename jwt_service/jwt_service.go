package jwt_service

import(
	"github.com/dgrijalva/jwt-go"
	"log"
)


type Claims struct {
	Password       string
	Username string
	jwt.StandardClaims
}

func GenerateToken(cla Claims) (string,error){
	token,err :=  jwt.NewWithClaims(jwt.SigningMethodHS256, cla).SignedString([]byte("golang"))
	if err != nil{
		log.Println("generate token falied:",err)
		return "",err
	}
	return token,nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	  return []byte("golang"), nil
	})
	if err != nil {
	  return nil, err
	}
	if tokenClaims != nil {
	  if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	  }
	}
	return nil, err
  }