package terminal_server

import(
	"net"
	"log"
	"Global/models"
	// "encoding/json"
	"bufio"
	"io"
	"strings"
	"strconv"
)

type Humidifier struct{
	Tem float32	//温度	
	Fog float32 //雾气浓度
	Water float32 //湿度
	Co2 float32 
	Co float32
	N2 float32
	No float32
	N2o float32
}

func Service(conn net.Conn){
	for{
		br := bufio.NewReader(conn)
		receive,err := br.ReadString('\n')
		if err == io.EOF{
			log.Println("read finish")
			break
		}
		if err != nil && err != io.EOF{
			log.Println("terminal server read string error:",err)
			break
		}
		receiveStr := strings.Split(receive, " ")
		receiveStr = receiveStr[:len(receiveStr)-1]
		if len(receiveStr) != 8{
			log.Println("terminal server reiceve nums failed:",len(receiveStr))
		}
		m := &models.HumidifierTable{
			Tem:stringToFloat32(receiveStr[0]),
			Fog :stringToFloat32(receiveStr[1]),
			Water:stringToFloat32(receiveStr[2]),
			Co2:stringToFloat32(receiveStr[3]),
			Co:stringToFloat32(receiveStr[4]),
			N2:stringToFloat32(receiveStr[5]),
			No:stringToFloat32(receiveStr[6]),
			N2o:stringToFloat32(receiveStr[7]),
		}
		err = models.HumidifierCreateData(m)
		if err != nil{
			log.Println("terminal server create data failed:",err)
			break
		}
	}
}

func stringToFloat32(value string) float32{
	valueF,_ := strconv.ParseFloat(value,32)
	valueF32 := float32(valueF)
	return valueF32
}

