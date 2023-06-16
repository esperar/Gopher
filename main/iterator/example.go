package main

import "fmt"

type Obj struct {
	Name string
	Age  int
}

func PrintObject(list []Obj) {
	for index, object := range list {
		fmt.Printf("index: %d, object: %+v\n", index, object)
	}
}

func main() {

	for index := 1; index <= 10; index++ {
		fmt.Println("index: %d\n", index)
	}

	for index := 1; index <= 10; {
		fmt.Println("index: %d\n", index)
		index++
	}

	index := 1
	for index <= 10 {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index = 1
	for true {
		fmt.Println("index: %d\n", index)
		index++

		if index > 10 {
			break
		}
	}

	// for range
	strArray := []string{"a", "b", "c"}
	for index, str := range strArray {
		fmt.Printf("index: %d, str: %s\n", index, str)
	}

	dictionary := map[string]string{
		"keyA": "valueA",
		"keyB": "valueB",
	}

	for key, value := range dictionary {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	list := []Obj{
		{"Beckham", 11},
		{"Zidane", 7},
		{"Ronaldo", 9},
	}

	for _, object := range list {
		object.Age = object.Age * 2
	}

	// 출력, 11, 7, 9
	PrintObject(list)

	for index := range list {
		object := &list[index]
		object.Age = object.Age * 2
	}

	// 출력, 22, 14, 18
	PrintObject(list)

}
