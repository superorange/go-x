package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func arrayTest(myArray *[4]int) {
	(*myArray)[0] = 100
}
func charT() {
	//a := 'A'
	//var az [26]byte
	//az[0] = 'A'
	//for i := 0; i < len(az); i++ {
	//	az[i] = az[i] + 1
	//	fmt.Printf("%c", az[i])
	//}
	var si []byte
	fmt.Printf("%T\n", si)
	//fmt.Printf("%T\n", a)

}

func newT() {
	a := new(interface{})
	*a = "80"
	fmt.Println(*a)
}
func copyT() {
	to := make([]int, 5)
	from := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(to, from)
	to[0] = 100
	fmt.Println(to)
	fmt.Println(from)
	to = append(to, 100)
	fmt.Println(to)
	fmt.Println("len:", len(to))

	to = to[1:1]
	to = to[0:0]
	fmt.Println("to:", to)
	fmt.Println("cap:", cap(to))
	fmt.Println("len:", len(to))
}

func strText() {
	a := "我爱你中国"
	runeToFind := '国'

	runes := []rune(a)
	for i, r := range runes {
		if r == runeToFind {
			fmt.Println(i) // 输出：3
			break
		}
	}

	// 或者使用 strings.IndexRune
	index := strings.IndexRune(a, runeToFind)
	fmt.Println(index) // 输出：3

	// 如果希望获得基于字符的索引
	charIndex := utf8.RuneCountInString(a[:index])
	fmt.Println(charIndex) // 输出：2
	//b := []byte(a)
	//for i, i2 := range b {
	//
	//	fmt.Printf("index:%d--value:%c\n", i, i2)
	//}
	//for len(b) > 0 {
	//	r, size := utf8.DecodeRune(b)
	//	fmt.Printf("size: %d,r:%c\n", size, r)
	//	b = b[size:]
	//}
	//
	//fmt.Println(b)
	//s, _ := utf8.de(b)
	//fmt.Println("s:", s)
}
func ioT() {
	//os.ReadFile()
}

func it() {
	s1 := make([]int, 0, 2)
	fmt.Println(s1)
	s2 := append(s1[:cap(s1)], 9)
	fmt.Println(s2)
	fmt.Println("cap:", cap(s2))
	fmt.Println("len:", len(s2))

}
func main() {
	it()
	//数组的3种初始化方式
	//array1 := [10]int{1, 2, 3}
	//array2 := [...]int{1, 2, 3}
	//var array3 = [10]int{1, 2, 3}
	//var array4 [10]int
	//array4[2] = 9
	//fmt.Println(array4)
	//var array5 = array4[:]
	//fmt.Printf("type:%T\n", array4)
	//fmt.Printf("type:%T\n", array5)
	//fmt.Printf("array1:%v\n", array1)
	//fmt.Printf("array2:%v\n", array2)
	//fmt.Printf("array3:%v\n", array3)
	//fmt.Printf("len %d\n", len(array2))

	//for i := 0; i < len(array1); i++ {
	//	fmt.Println(array1[i])
	//}
	//for i, i2 := range array1 {
	//	fmt.Println(i, i2)
	//}

	//array1 := [...]int{1, 2, 3, 4}
	////arrayTest(array1)
	//fmt.Println(array1)
	//arrayTest(&array1)
	//fmt.Println(array1)
}
