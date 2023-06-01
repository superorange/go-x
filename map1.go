package main

import (
	"fmt"
)

// 切片和map一样，都是引用类型
// 使用 make来分配切片和map
func main() {
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["cnm"] = 20
	map1["cnm1"] = 20
	fmt.Println(map1)
	//err := json.Encoder.Encode(map1)
	//if err != nil {
	//	return
	//}

	map2 := make(map[string]interface{})
	map2["ppp"] = 12
	map2["ppp1"] = "zjangs"
	map2["q"] = []int{1, 2, 3, 4}
	fmt.Println(map2)

	m3 := map[string]func(a int) int{
		"b": func(a int) int {
			return a + a
		},
	}

	m3["a"] = func(a int) int {
		return a
	}

	fmt.Printf("%s\n", m3["a"])

	if c, ok := map2["ppp"]; ok {
		fmt.Println(c)
		fmt.Println(ok)
	}
	delete(map2, "ppp")

	for k, v := range map2 {
		fmt.Println(k, v)
	}

}
