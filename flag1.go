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

func t19() {
	file, err := os.OpenFile("t.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error open: ", err)
		return
	}
	writeBuffer := bufio.NewWriter(file)

	writeBuffer.Write([]byte("hello world\n"))
	writeBuffer.Flush()

}

func t20() {
	listA := make([]int, 10)
	listA[1] = 99
	listB := listA[:]
	listB[9] = 90
	var listC []int
	copy(listC, listB)
	listC[2] = 88
	fmt.Println(listA)
	fmt.Println(listB)
	fmt.Println(listC)
}

func t21() {
	listA := make([]int, 10)
	listA[1] = 99
	listB := make([]int, len(listA))
	copy(listB, listA)
	listB[2] = 90
	fmt.Println(listA)
	fmt.Println(listB)

	listC := append(listA)
	listC[3] = 88
	fmt.Println(listC)

	fmt.Printf("listA len:%d ,cap:%d\n", len(listA), cap(listA))
	fmt.Printf("listB len:%d ,cap:%d\n", len(listB), cap(listB))
	fmt.Printf("listC len:%d ,cap:%d\n", len(listC), cap(listC))

}

func main() {
	t21()
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
