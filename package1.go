package main

import (
	"container/list"
	"fmt"
)

// /数组切片使用 make
// /结构体使用new
// /在一个结构体中对于每一种数据类型只能有一个匿名字段
func init() {

	fmt.Println("init")
}

func ee() {
	//syscall.Syscall(syscall.SYS_REBOOT, syscall.LINUX_REBOOT_MAGIC1, syscall.LINUX_REBOOT_MAGIC2, syscall.LINUX_REBOOT_CMD_RESTART)
	//items := []string{"a", "b", "c", "d", "e", "f"}
	//itemP := &items
	//for e := itemP.Front(); e != nil; e = e.Next() {
	//	println(e.Value)
	//}
	l := list.New()

	for e := l.Front(); e != nil; e.Next() {
		fmt.Println(e)
	}
	fmt.Println(l.Front())

}
