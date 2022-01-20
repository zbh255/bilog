package bilog

import "strconv"

// 快速格式化年月日

var yearBuf = func() []string {
	strs := make([]string, 4096-1970)
	for k := range strs {
		strs[k] = strconv.Itoa(1970 + k)
	}
	return strs
}()

var monthBuf = func() []string {
	strs := make([]string, 11)
	for k := range strs {
		strs[k] = strconv.Itoa(1 + k)
	}
	return strs
}()

var dayBuf = func() []string {
	strs := make([]string, 31)
	for k := range strs {
		strs[k] = strconv.Itoa(1 + k)
	}
	return strs
}()

func fastConvertYear(i int) string {
	return yearBuf[i-1970]
}

func fastConvertMonth(i int) string {
	return monthBuf[i-1]
}

func fastConvertDay(i int) string {
	return dayBuf[i-1]
}
