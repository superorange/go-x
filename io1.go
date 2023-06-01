// 从控制台读取输入:
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func t1() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s -- %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
func t2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\a')
	if err != nil {
		fmt.Printf("Found an error: %s\n", err)
	} else {
		fmt.Printf("Your input is %s---\n", input)
	}
	writeString, err := os.Stdout.WriteString("This is a test\n")
	if err != nil {
		return
	}

	fmt.Printf("The length of string is %d\n", writeString)
}
func t3() {
	inputFile, err := os.Open("io1.go")
	if err != nil {
		return
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Println("Error close: ", err)
			return
		}
	}(inputFile)
	//func byteRead(){
	//	r := 1
	//	for r > 0 {
	//		buf := make([]byte, 1024)
	//		r, err = inputFile.Read(buf)
	//		if err != nil {
	//			log.Println("Error read: ", err)
	//			return
	//		}
	//		fmt.Printf("Received data: %v", string(buf[:r]))
	//	}
	//}

	//bufio.NewWriter()
	fileReader := bufio.NewReader(inputFile)
	for {
		readString, err := fileReader.ReadString('\r')
		if err != nil {
			//byteRead()
			log.Println("Error read: ", err)
			return
		}
		fmt.Printf("The first line is %s\n", readString)
	}

}

func t4() {
	file, err := os.ReadFile("io1.go")
	if err != nil {
		log.Println("Error read: ", err)
		return
	}
	fmt.Printf("The first line is %s\n", file)

}

type Book struct {
	title    string
	price    float64
	quantity int
}

func (book Book) String() string {
	return fmt.Sprintf("Book{title: %s, price: %f, quantity: %d}\n", book.title, book.price, book.quantity)

}

func t5() {
	file, err := os.ReadFile("t.txt")
	if err != nil {
		log.Println("Error read: ", err)
		return
	}
	s2 := string(file)
	split := strings.Split(s2, "\n")
	split = split[:len(split)-1]
	tag := make([]Book, 0)
	for _, s3 := range split {
		i2 := strings.Split(s3, ";")
		book := new(Book)
		book.title = i2[0]
		book.price, _ = strconv.ParseFloat(i2[1], 64)
		book.quantity, _ = strconv.Atoi(i2[2])
		tag = append(tag, *book)
	}
	fmt.Println(tag)

}
func t6() {
	s := "3.1415926535" // 输入字符串
	precision := 4      // 小数位数精度

	// 格式化输入字符串，限制小数位数
	formatted := fmt.Sprintf("%.*f", precision, s)

	// 将格式化后的字符串转换为浮点数
	f, err := strconv.ParseFloat(formatted, 64)
	if err != nil {
		fmt.Println("Failed to parse float:", err)
		return
	}

	fmt.Println("Parsed float:", f)
}

func t7() {
	//err := os.WriteFile("t.txt", []byte("The Go Programming Language\n"), 0644|fs.ModeAppend)
	//if err != nil {
	//	log.Println("Error write: ", err)
	//	return
	//}
	//file, _ := os.ReadFile("t.txt")
	//fmt.Println(string(file))
	file, _ := os.OpenFile("t.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error close: ", err)
			return
		}
	}(file)
	_, err := file.WriteString("The Go Programming Language\n")
	if err != nil {
		log.Println("Error write: ", err)
		return
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}
	all, _ := io.ReadAll(file)
	fmt.Println(string(all))

}
func main() {
	t7()
	//fmt.Println(os.Args)
	//t6()
	// 输出结果: From the string we read: 56.12 5212 Go
}
