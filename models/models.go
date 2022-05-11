package models

//操作数据库
import(
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
	"log"
	"errors"
	"context"
)

var DB *gorm.DB
var RDB *redis.Client

type UserTable struct{
	gorm.Model
	UserName string	`gorm:"type:varchar(255)"`
	Password string
}

type HumidifierTable struct{
	gorm.Model
	Tem float32	//温度
	Fog float32 //雾气浓度
	Water float32 //湿度
	Co2 float32 
	Co float32
	N2 float32
	No float32
	N2o float32
}

//查询最新的数据

//create 
func UserCreateData(name,password string) error{
	err := DB.AutoMigrate(&UserTable{})
	if err != nil{
		log.Println("create data:auto migreate failed:",err)
		return err
	}
	u := &UserTable{
		UserName:name,
		Password:password,
	}
	//update in redis
	var ctx = context.Background()
	rdbErr := RDB.Set(ctx, name, password, 0)
	if err != nil{
		log.Println("redis update failed:", rdbErr.Err())
		return rdbErr.Err()
	}
	return DB.Create(u).Error
}

//find
func UserFindData(name string) (*UserTable,error){
	//redis + mysql
	var cxt = context.Background()
	u := &UserTable{}
	password,err := RDB.Get(cxt, name).Result()
	if err == redis.Nil{
		if err := DB.Where("user_name = ?", name).Find(u).Error; err != nil{
			log.Println("user find data:find falied:",err)
			return u,err
		}
		//update in redis
		rdbErr := RDB.Set(cxt, u.UserName, u.Password, 0)
		if rdbErr.Err() != nil{
			log.Println("redis update failed:", rdbErr.Err())
			return u,rdbErr.Err()
		}
		return u,nil
	}else if err != nil{
		return u,err
	}else{
		u.UserName = name
		u.Password = password
		return u,nil
	}
}

func HumidifierCreateData(h *HumidifierTable) error{	
	err := DB.AutoMigrate(&HumidifierTable{})
	if err != nil{
		log.Println("create data: auto migrate error:",err)
		return err
	}
	return DB.Create(h).Error
}

func HumidifierGetData()(*HumidifierTable, error){
	h := &HumidifierTable{}
	result := DB.Last(h)
	if errors.Is(result.Error, gorm.ErrRecordNotFound){
		log.Println("humidifier table if empty")
		return h,errors.New("tableEmpty")
	}
	if result.Error != nil{
		log.Println("humidifier get data failed:",result.Error)
		return h, result.Error
	}
	return h,nil
}

