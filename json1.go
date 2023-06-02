package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
}

type PersonModel struct {
	Person *Person `json:"Person"`
}

func t1() {
	open, err := os.Open("test.json")
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Println("Error close: ", err)
		}
	}(open)
	if err != nil {
		return
	}
	decoder := json.NewDecoder(open)
	var person Person
	err = decoder.Decode(&person)
	if err != nil {
		log.Println("Error decode: ", err)
		return
	}
	fmt.Printf(person.LastName)
	fmt.Println(person)
	file, _ := os.ReadFile("test.json")

	data := make(map[string]interface{})
	err = json.Unmarshal(file, &data)
	fmt.Println(data["Person"].(map[string]interface{})["FirstName"])
}

func t2() {
	file, _ := os.Open("test.json")
	decoder := json.NewDecoder(file)
	var model PersonModel
	decoder.Decode(&model)
	fmt.Println(model)
}

func t4() {
	jsonStr := "{\n  \"Person\": [\n    {\n      \"FirstName\": \"Laura\",\n      \"LastName\": \"Lynn\"\n    },\n    {\n      \"FirstName\": \"Laura\",\n      \"LastName\": \"Lynn\"\n    }\n  ]\n}\n"
	var maps map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &maps)

	if person, ok := maps["Person"].(map[string]interface{}); ok {
		fmt.Println("person is of type map[string]interface{}", person)
	}

	switch v := maps["Person"].(type) {

	case map[string]interface{}:
		fmt.Println("person is of type map[string]interface{}", v)
	case []interface{}:
		fmt.Println("person is of type []interface{}", v)
	default:
		fmt.Println("I don't know about type %T!\n", v)

	}

}

func t5() {
	jsonStr := "{\n  \"Person\": [\n    {\n      \"FirstName\": \"Laura\",\n      \"LastName\": \"Lynn\"\n    },\n    {\n      \"FirstName\": \"Laura\",\n      \"LastName\": \"Lynn\"\n    }\n  ]\n}\n"
	var maps map[string][]map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &maps)

	fmt.Println(maps["Person"][0]["FirstName"])
	file, _ := os.OpenFile("test1.json", os.O_CREATE|os.O_WRONLY, 0666)

	marshal, _ := json.Marshal(maps)

	file.Write(marshal)

}

func t6() {
	file, _ := os.OpenFile("test2.json", os.O_CREATE|os.O_WRONLY, 0666)

	person := PersonModel{Person{"张三", "李四"}}

	marshal, _ := json.Marshal(person)
	file.Write(marshal)
	file.Close()

}

func main() {
	t6()
}
