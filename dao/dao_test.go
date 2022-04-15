package dao

import(
	"testing"
	"Global/models"
	"fmt"
	"gorm.io/gorm"
	"errors"
)

func TestModelFind(t *testing.T){
	err := MySQLInit()
	if err != nil{
		t.Errorf("failed:%#v",err)
	} 
	db,err := models.UserFindData("xiaoming")
	models.UserCreateData("lll","123")
	if db.UserName == ""{
		fmt.Println("no this name")
	}else{
		fmt.Println(db.UserName,db.Password)
	}
	if err != nil{
		t.Errorf("find failed%#v",err)
	}
}

func TestModelHumFind(t *testing.T){
	MySQLInit()
	h := &models.HumidifierTable{
		Tem:12.1,
		Co2:18.90,
		N2:6,
	}
	err := models.HumidifierCreateData(h)
	if err != nil{
		t.Errorf("error")
	}
	result,err := models.HumidifierGetData()
	if errors.Is(err, gorm.ErrRecordNotFound){
		fmt.Println("ture")
	}else{
		fmt.Println("false")
	}
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			fmt.Println("no this data")
		}else{
			t.Errorf("error:%#v",err)
		}
	}else{
		fmt.Println("get data :",result)
	}
}