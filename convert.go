package bilog

import (
	"reflect"
	"strconv"
	"unsafe"
)

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
	strs := make([]string, 12)
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
		strs[k] = strconv.Itoa(k) + ":"
		if k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

// 分钟有00，索引k无需+1
var minuteBuf = func() []string {
	strs := make([]string, 60)
	for k := range strs {
		strs[k] = strconv.Itoa(k) + ":"
		if k < 10 {
			strs[k] = "0" + strs[k]
		}
	}
	return strs
}()

// 秒钟有00, 索引k无需+1
var secondBuf = func() []string {
	strs := make([]string, 60)
	for k := range strs {
		strs[k] = strconv.Itoa(k) + " "
		if k < 10 {
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

// 小时有00，寻址时不需要-1
func fastConvertHour(i int) string {
	return hourBuf[i]
}

// 分钟有00，寻址时不需要-1
func fastConvertMinute(i int) string {
	return minuteBuf[i]
}

// 秒钟有00，寻址时不需要-1
func fastConvertSecond(i int) string {
	return secondBuf[i]
}

// 返回存储时间的字节数组和写入数据的有效数量
func fastConvertAllToArray(year, month, day, hour, minute, second int) (tmp [32]byte, eff int) {
	data := (uintptr)(unsafe.Pointer(&tmp))
	slice := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: data,
		Len:  0,
		Cap:  len(tmp),
	}))
	slice = append(slice, fastConvertYear(year)...)
	slice = append(slice, fastConvertMonth(month)...)
	slice = append(slice, fastConvertDay(day)...)
	slice = append(slice, fastConvertHour(hour)...)
	slice = append(slice, fastConvertMinute(minute)...)
	slice = append(slice, fastConvertSecond(second)...)
	return tmp, len(slice)
}

func fastConvertAllToSlice(year, month, day, hour, minute, second int) []byte {
	tmp := make([]byte, 0, 32)
	tmp = append(tmp, fastConvertYear(year)...)
	tmp = append(tmp, fastConvertMonth(month)...)
	tmp = append(tmp, fastConvertDay(day)...)
	tmp = append(tmp, fastConvertHour(hour)...)
	tmp = append(tmp, fastConvertMinute(minute)...)
	tmp = append(tmp, fastConvertSecond(second)...)
	return tmp
}
