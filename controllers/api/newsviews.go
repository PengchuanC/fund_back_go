package api

import (
	"fund_back_go/databases"
	"fund_back_go/models"
	"fund_back_go/util"
	"github.com/gin-gonic/gin"
)

func NewsViews(c *gin.Context){
	var news models.News
	var ret []models.News
	var count, pageNumber int64

	page := util.ConvertStr2Int(c.Query("page"))
	pageNumber = 25
	databases.DB.Model(&news).Count(&count)

	databases.DB.Offset(pageNumber*page).Find(&ret).Limit(pageNumber)
	c.JSON(200, gin.H{
		"total": count, "per_page": pageNumber, "data": ret,
	})
}
