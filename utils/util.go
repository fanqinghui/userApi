package utils

import (
	"fmt"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func DoJsonData(body []byte) string {
	bodyStr := string(body)
	bodyStr = Substr(bodyStr, 1, strings.LastIndex(bodyStr, "'")-1)
	bodyStr = strings.Replace(bodyStr, "{", "{\"", 1)
	fmt.Println("body:" + bodyStr)
	bodyStr = strings.Replace(bodyStr, ":", "\":\"", 1)
	bodyStr = strings.Replace(bodyStr, "}", "\"}", 1)
	return bodyStr
}
