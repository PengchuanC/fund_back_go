package main

import (
	"fund_back_go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	err := routes.Engine.Run(":5000")
	if err != nil {
		recover()
	}
}
