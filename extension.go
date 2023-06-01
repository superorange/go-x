package main

import "fmt"

type mapExtension map[string]interface{}

func (m mapExtension) hasElement(e string) bool {
	_, ok := m[e]
	return ok

}

type mapStruct struct {
	data map[string]interface{}
}

func (m *mapStruct) hasElement(e string) bool {
	_, ok := m.data[e]
	return ok

}

func main() {
	//mp1 := mapExtension{"a": 1, "b": 2, "c": 3}
	//fmt.Println(mp1.hasElement("t"))
	data := map[string]interface{}{"a": 1, "b": 2, "c": 3}
	mp1 := mapStruct{data: data}
	fmt.Println(mp1.hasElement("t"))

}
