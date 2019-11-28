package api

import (
	"fund_back_go/databases"
	"fund_back_go/models"
	"fund_back_go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
**处理新闻请求，一般附带4个参数
**@param page: 请求的页码，与数据库分页对应
**@param section: 请求的板块，一般分为全部、宏观、金融...
**@param date: 请求的日期
**@param search: 请求附带关键字，按关键字检索
 */
func NewsViews(c *gin.Context) {
	var news models.News
	var ret []models.News
	var count, pageNumber int64

	page := util.ConvertStr2Int(c.DefaultQuery("page", "1"))
	pageNumber = 25
	databases.DB.Model(&news).Count(&count)

	databases.DB.Offset(pageNumber * page).Limit(pageNumber).Find(&ret)
	c.JSON(200, gin.H{
		"total": count, "per_page": pageNumber, "data": ret,
	})
}

/*
**Method: GET
**Route: /newslist
 */
func NewsList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	section := c.Query("section")
	date := c.Query("date")
	search := c.Query("search")

	var news models.News
	var ret []models.News
	db := databases.DB

	if section != "" {
		db = db.Where("keyword = ?", section)
	}
	if date != "" {
		db = db.Where("savedate like ?", date+"%")
	}
	if search != "" {
		db = db.Where("title like ?", "%"+search+"%").Or("abstract like ?", "%"+search+"%")
	}

	var _page, count, pageNumber int64
	pageNumber = 25
	_page = util.ConvertStr2Int(page)
	db.Model(&news).Count(&count)
	db.Offset(pageNumber * _page).Limit(pageNumber).Find(&ret)
	c.JSON(http.StatusOK, gin.H{
		"total": count, "per_page": pageNumber, "data": ret,
	})
}
