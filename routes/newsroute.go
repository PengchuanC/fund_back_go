package routes

import "fund_back_go/controllers/api"

func init() {
	news := v1.Group("/news")
	{
		news.GET("/", api.NewsViews)
		news.GET("/newslist", api.NewsList)
		news.POST("/follow", api.Follow)
		news.Any("/follow/keywords", api.FollowKeyword)
	}
}
