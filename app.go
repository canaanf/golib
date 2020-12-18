package lib

import (
	"log"
)

// 对错误的一般处理，打印并退出应用程序
func ExecErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}