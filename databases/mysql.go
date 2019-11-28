package databases

import (
	"fmt"
	"fund_back_go/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	var mysql conf.DatabaseConfig
	var mode, uri string

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_ff_" + defaultTableName
	}

	mode = conf.Config.Mode
	switch mode {
	case "development":
		mysql = conf.Config.Development
	case "production":
		mysql = conf.Config.Production
	default:
		mysql = conf.Config.Development
	}

	uri = fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.Database,
	)
	DB, err = gorm.Open("mysql", uri)

	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v\n", DB.Error)
	}

	DB.SingularTable(true)
}
