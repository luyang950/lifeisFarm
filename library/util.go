package library

import (
	"strconv"
)

var yearMonthDay = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 21,
}

func YearDay(month, day int) int {
	var dayNum int

	switch month {
	case 1:
		dayNum = 0
	case 2:
		dayNum = 31
	case 3:
		dayNum = 59
	case 4:
		dayNum = 90
	case 5:
		dayNum = 120
	case 6:
		dayNum = 151
	case 7:
		dayNum = 181
	case 8:
		dayNum = 212
	case 9:
		dayNum = 243
	case 10:
		dayNum = 273
	case 11:
		dayNum = 304
	case 12:
		dayNum = 334
	default:
		return -1
	}

	dayNum += day

	// 配置文件中从0开始
	dayNum--

	return dayNum
}

func YearDay2Date(yearDay int) string {
	var date string
	var month, day = 1, 0

	for i := 1; i < 13; i++ {
		if yearDay-yearMonthDay[i] > 0 {
			month++
			yearDay -= yearMonthDay[i]
			continue
		}

		day = yearDay
	}

	date = strconv.Itoa(month) + "-" + strconv.Itoa(day)

	return date
}
