package router

import(
	"testing"
)

func htmlTest(t *testing.T){
		r := RountersInit()
		r.Run("127.0.0.1:9527")
	
}