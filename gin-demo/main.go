package main

import "github.com/gin-gonic/gin"

// curl http://localhost:8080/
func main() {
	engine := gin.New()
	engine.GET("/", Get)
	engine.POST("/testShouldBind", TestShouldBind)
	engine.GET("/testRetry", TestRetry)
	engine.POST("/importData", ImportCheckinQuestion)
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
