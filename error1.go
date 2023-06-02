package main

import (
	"errors"
	"fmt"
)

var me error = errors.New("go出错啦")

func t1() {
	panic(me)
}
func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("捕获到异常:%v\n", err)
		}
	}()

	fmt.Printf("error:%v\n", me)
	t1()
	println("continue")
}
