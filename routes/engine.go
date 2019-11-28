package routes

import "github.com/gin-gonic/gin"

var Engine *gin.Engine
var v1 *gin.RouterGroup

func init() {
	Engine = gin.Default()
	v1 = Engine.Group("/api/v1")
}
