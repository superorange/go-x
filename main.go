package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type CNM int
type okex func(a int, b int) int

//const timeFormat = "2006 01 02 15:04:05"
const (
	yyyy = "2006"
	MM   = "01"
	DD   = "02"
	HH   = "15"
	mm   = "04"
	ss   = "05"
)

//type FS = func()
func init1() {
	var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}
func test1() {
	a := 10
	abs := math.Abs(float64(a))
	fmt.Println("abs:", abs)
}
func test2() {
	str := "10"
	st, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	openFile, err := os.OpenFile("./main.go", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	b2 := make([]byte, 30)
	_, err = openFile.Read(b2)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("b2:", string(b2))
	fmt.Println("st:", st)

}
func test3() {
	str := "1998"
	switch str {
	case "90":
		fmt.Println("90")
	case "78", "1998":
		fmt.Println("78")
	default:
		fmt.Println("default")

	}

}
func test4() {
	var str = "GGGGGG"
	for i := 0; i < len(str); i++ {
		println(strings.Repeat("G", i))
	}

}
func test5() {
	var h = 20
	var w = 20
	for i := 0; i < h; i++ {
		for j := 0; j < w-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	var i = 10
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}

}

func test6() *int {
	result := 7 * 8 * 10
	return &result
}
func test7(a *int) {
	*a = 10
}
func test8() (a int, b int) {
	a = 6
	b = 8
	return
}
func test9(items ...interface{}) {
	for i, value := range items {

		fmt.Println("value:", value)
		fmt.Println("i:", i)
	}
}

func deferTest(a int) (b int) {
	fmt.Println("run over,time is ", time.Now().Format("2006-01-02 15:04:05"))
	return b
}
func start() (a int) {
	fmt.Println("run start,time is ", time.Now().Format("2006-01-02 15:04:05"))
	return
}

func def2(str string) (s string, err error) {
	defer func() {
		log.Print("函数运行结束:", str, s, err)
	}()
	return s, err
}

type myFunc func(a int, b int)

func test10(a int, myFunc2 myFunc) {
	myFunc2(a, 10)
}
func test11() {
	myFunc2 := func(a int, b int) {
		fmt.Println("hello world")
	}
	test10(10, myFunc2)
}

func f() (ret int) {
	defer func() {
		fmt.Println("ret:", ret)
		ret++
	}()
	return 10
}
func test12() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()
	println("2023/05/30 16:37:27 /Users/tian/projects/private/go/go-study/main.go:85\n")
	log.Print()

}

func test13() {
	log.SetPrefix("main.go--:")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("")
	var where = log.Print
	where()
	where()
	where()
}

func test14() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %d\n", delta.Milliseconds())
}
func main() {
	items := []int{1, 2, 3, 4}[:]
	items2 := make([]int, 10, 10)
	fmt.Printf("type:%T\n", items)
	fmt.Println(items)
	fmt.Printf("type2:%T\n", items2)
	//items2[49] = 89
	items2 = append(items2, 100)
	items2 = append(items2, 101)
	fmt.Println("cap:", cap(items2))
	fmt.Println("len:", len(items2))
	fmt.Println(items2)

	//test14()
	//func() {
	//	log.Print("函数运行开始:")
	//}()
	//
	//if r, err := def2("80"); err == nil {
	//	log.Printf("main: %s -- %s\n", r, err)
	//}
	//defer deferTest(start())
	//time.Sleep(10 * time.Second)
	//fmt.Println("run main,time is ", time.Now().Format("2006-01-02 15:04:05"))
	//defer fmt.Printf("%d ", 1)
	//defer fmt.Printf("%d ", 2)
	//defer fmt.Printf("%d ", 3)
	//test9(1, "张三", 'c', "李四")

	//i2, _ := test8()
	//_, i3 := test8()
	//fmt.Println("i2:", i2)
	//fmt.Println("i3:", i3)
	////i2 := test6()
	////fmt.Println("i2:", *i2)
	//
	//z := 20
	//test7(&z)
	//fmt.Println("a:", z)

	return

	str := "我是DART语言开发教程"
	ptr1 := &str
	ptr2 := &ptr1
	ptr3 := &ptr2
	fmt.Println("str:", str)
	fmt.Println("ptr1 v:", *ptr1)
	fmt.Println("ptr2 v:", *ptr2)
	fmt.Println("ptr3 v:", **ptr3)
	fmt.Println("ptr1 address:", &ptr1)
	fmt.Println("ptr2 address:", &ptr2)
	fmt.Println("ptr3 address:", &ptr3)
	if str == "" {
		fmt.Println("str is empty")
	} else if str == "x" {

	} else {

	}

	//a := 10
	//if a > 10 {
	//	fmt.Println("a>10")
	//} else if a > 11 {
	//
	//} else {
	//
	//}

	//str := "你好中国"
	//ptr := &str
	//*ptr = "CNM"
	//fmt.Println("str:", str)
	//fmt.Println("ptr:", *ptr)
	//fmt.Println("ptr address:", &ptr)
	//fmt.Printf("ptr address :%p\n", &ptr)
	//
	//now := time.Now()
	//fmt.Println("now:", now.Format(yyyy+"-"+MM+"-"+DD+"-"+HH+mm+ss))
	return
	//str := "我是DART语言开发教程"
	//reader := strings.NewReader(str)
	////bytes.NewBufferString(str)
	//b2 := make([]byte, 30)
	//read, err := reader.Read(b2)
	//if err != nil {
	//	return
	//}
	//fmt.Println("read:", read)
	//fmt.Println("b2:", string(b2))
	//return
	//strings.HasPrefix("abc", "a")
	//strings.Split("a,b,c", ",")
	//str := "你 好      吗 你"
	//fmt.Println("index:", strings.IndexRune(str, rune('吗')))
	//replace := strings.Replace(str, "你", "我", -1)
	//fmt.Println("replace:", replace)
	//fmt.Println("str:", str)
	//fmt.Println("count:", strings.Count(str, "你"))
	//fmt.Println("repeat:", strings.TrimLeft(str, "你"))
	//fmt.Println("split Fields", len(strings.Fields(str)))
	//fmt.Println("split", len(strings.Split(str, " ")))
	//fmt.Println("split regexp", len(re.MustCompile("\\s*").Split(str, -1)))
	//return
	//const a, b, c = 1, 2, 3
	//const d, e, f = true, 0.5, "Hello"
	//const (
	//	Male   = 0
	//	Female = 1
	//)
	//// 因式分解法 全局变量
	//var (
	//	age    int
	//	height float32
	//)
	//g:=10
	//
	//println(age)
	//println(height)
	//println(Male)
	//
	//fmt.Println("Hello, World!")
	goos := runtime.GOOS
	fmt.Printf("system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	//var a string = "abc"
	//fmt.Println("hello, world")
	var a, b, c = 10, "50", 99
	fmt.Println(a, b, c)
	d, e, f := 10, "50", 99
	d, f = f, d
	fmt.Println(d, e, f)
	i := rand.Int()
	if i != 10 {
		fmt.Println(i)
	}
	for i := 0; i < 20; i++ {
		println("I:", i)
	}
	var cnm CNM = 89
	fmt.Println(cnm)
}
