package main

import "fmt"

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

}
