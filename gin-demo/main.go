package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// curl http://localhost:8080/
func main() {
	fmt.Println("1234")
	engine := gin.New()
	engine.GET("/", Get)
	engine.POST("/testShouldBind", TestShouldBind)
	engine.GET("/testRetry", TestRetry)
	engine.POST("/importData", ImportCheckinQuestion)
	engine.GET("/testAfter", TestAfter)
	engine.GET("/testGo", TestMutiGo)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
