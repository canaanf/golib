// Package funcs 为canaanf提供常用工具函数包
package lib

import (
	"fmt"
	"unsafe"
)

// 对参数in数据的最低几个字节（长度由参数size给出）反转字节序，指定size超过sizeof(uint64)时返回错误
func ReverseIntBytes(in uint64, size int) (out uint64, err error) {
	if size > int(unsafe.Sizeof(out)) {
		err = fmt.Errorf("Size must less than %d.", unsafe.Sizeof(out))
		return
	}

	for i := 0; i < size; i++ {
		out <<= 8
		out |= in & 0xff
		in >>= 8
	}
	return
}