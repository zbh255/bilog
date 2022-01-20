package bilog

import "strconv"

// 快速格式化年月日
// year之外的数字如果小于10则会使用零填充prefix

var yearBuf = func() []string {
	strs := make([]string, 4096-1970)
	for k := range strs {
		strs[k] = strconv.Itoa(1970+k) + "-"
	}
	return strs
}()

var monthBuf = func() []string {
	strs := make([]string, 11)
	for k := range strs {
		strs[k] = strconv.Itoa(1+k) + "-"
		if 1+k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

var dayBuf = func() []string {
	strs := make([]string, 31)
	for k := range strs {
		strs[k] = strconv.Itoa(1+k) + " "
		if 1+k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

var hourBuf = func() []string {
	strs := make([]string, 24)
	for k := range strs {
		strs[k] = strconv.Itoa(1+k) + ":"
		if 1+k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

var minuteBuf = func() []string {
	strs := make([]string, 60)
	for k := range strs {
		strs[k] = strconv.Itoa(1+k) + ":"
		if 1+k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

var secondBuf = func() []string {
	strs := make([]string, 60)
	for k := range strs {
		strs[k] = strconv.Itoa(1+k) + " "
		if 1+k < 10 {
			strs[k] = "0" + strs[k]
		}
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

func fastConvertHour(i int) string {
	return hourBuf[i-1]
}

//TODO: 修复分钟和秒的00产生的索引越界，0分时索引为-1
func fastConvertMinute(i int) string {
	return minuteBuf[i-1]
}

func fastConvertSecond(i int) string {
	return secondBuf[i-1]
}
