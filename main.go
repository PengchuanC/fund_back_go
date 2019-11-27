package main

import (
	api "fund_back_go/controllers/api"
	"fund_back_go/databases"
	"fund_back_go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	test()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	/*新闻视图*/
	news := r.Group("/api/v1/news")
	{
		news.GET("/", api.NewsViews)
	}

	r.Any("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"password": test(),
		})
	})

	_ = r.Run(":5000")
}

func test() string {
	db := databases.DB
	var news models.News
	db.First(&news)
	return news.Title
}
