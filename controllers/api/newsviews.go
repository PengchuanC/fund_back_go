package api

import (
	"fund_back_go/database"
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

/*
**新闻首页
**Method: GET
**Route: /breaking
 */
func NewsViews(c *gin.Context) {
	var news models.News
	var ret []models.News
	var count, pageNumber int64

	page := util.ConvertStr2Int(c.DefaultQuery("page", "1"))
	pageNumber = 25
	database.DB.Model(&news).Count(&count)

	database.DB.Offset(pageNumber * page).Limit(pageNumber).Find(&ret)
	c.JSON(200, gin.H{
		"total": count, "per_page": pageNumber, "data": ret, "page": page,
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
	db := database.DB

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
		"total": count, "per_page": pageNumber, "data": ret, "page": _page,
	})
}

/*
**Method: POST
**Route: /follow
 */
func Follow(c *gin.Context) {
	keyword := c.Query("keyword")
	page := c.DefaultQuery("page", "1")

	_page := util.ConvertStr2Int(page)

	var news models.News
	var ret []models.News
	var pageNum, count int64
	pageNum = 25

	db := database.DB
	db = db.Where(models.News{Keyword: keyword})
	db.Model(&news).Count(&count)
	db.Offset(_page * pageNum).Limit(pageNum).Find(&ret)
	c.JSON(http.StatusOK, gin.H{
		"total": count, "per_page": pageNum, "data": ret, "page": _page,
	})
}

/*
**Method: POST
**Route: /follow/keywords
 */
func FollowKeyword(c *gin.Context) {
	type Keyword struct {
		Keyword string
	}
	var ret []Keyword
	db := database.DB
	db.Table("t_ff_news").Select("distinct(keyword) as keyword").Where("keyword is not null").Scan(&ret)

	var data []string
	for _, key := range ret {
		data = append(data, key.Keyword)
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
