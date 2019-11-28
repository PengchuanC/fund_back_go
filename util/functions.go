package util

import (
	"strconv"
	"time"
)

func ConvertStr2Int(str string) int64 {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return int64(num)
}

/*字符串转时间*/
func ConvertStr2Date(str string) time.Time {
	birthday := "2006-01-02"
	loc, err := time.LoadLocation("Local")
	if err != nil {
		recover()
	}
	date, err := time.ParseInLocation(birthday, str, loc)
	if err != nil {
		recover()
	}
	return date
}

/*时间转字符串*/
func ConvertDate2Str(date time.Time) string {
	birthday := "2006-01-02"
	str := date.Format(birthday)
	return str
}
