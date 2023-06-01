package main

import (
	"bufio"
	"flag" // command line option parser
	"fmt"
	"log"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func t17() {
	if flag.NArg() == 0 {
		log.Println("please input open file")
		return
	}

}

func t18() {
	open, err := os.Open("flag1.go")
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {

		}
	}(open)
	if err != nil {
		log.Println("Error open: ", err)
		return
	}
	reader := bufio.NewReader(open)
	for {
		rByte, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("read exit : ", err)
			return
		}
		fmt.Printf(string(rByte))

	}

}

func main() {
	t18()
	return
	// 定义命令行参数并关联变量
	name := flag.String("name", "张三", "你的名字")
	age := flag.Int("age", 0, "你的年龄")

	flag.Parse()

	fmt.Println("名字:", *name)
	fmt.Println("年龄:", *age)

	fmt.Println("参数帮助信息:")
	flag.PrintDefaults()
}
