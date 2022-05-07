package bilog

import "io"

// 该函数会panic的两种情况
// - io.Writer对应的实现返回error
// - 写入的长度与data的长度不相同
func writeHandle(w io.Writer,data []byte) {
	n, err := w.Write(data)
	if err != nil {
		panic(err)
	}
	if n != len(data) {
		panic("write byte not equal")
	}
}
