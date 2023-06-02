package main

import "fmt"

type Name struct {
	FirstName string
}

type Person struct {
	Name *Name
}

func (person Person) String() string {
	return fmt.Sprintf("{name:%s}", person.Name)
}
func main() {
	name := Name{FirstName: "张三"}

	n2 := new(Name)
	n2.FirstName = "张三"

	person := Person{&name}
	fmt.Println("name:", name)
	fmt.Println("person:", person)
	name.FirstName = "李四"
	fmt.Println("name:", name)
	fmt.Println("person:", person)

}
