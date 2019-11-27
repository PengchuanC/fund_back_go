package util

import "strconv"

func ConvertStr2Int(str string) int64 {
	num, err := strconv.Atoi(str)
	if err != nil{
		panic(err)
	}
	return int64(num)
}
