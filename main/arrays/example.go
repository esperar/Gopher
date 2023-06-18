package main

import "fmt"

func main() {

	// 배열 선언
	var intArray [3]int
	intArray[0] = 3
	intArray[1] = 4
	intArray[2] = 6

	// 선언과 동시에 초기화
	var stringArrayWithInit = [3]string{"A", "B", "C"}
	for index := range stringArrayWithInit {
		fmt.Println("%s\n", stringArrayWithInit[index])
	}

	// 멀티플 배열
	var intMultipleArray = [2][3]int{
		{3, 4, 5},
		{4, 5, 6},
	}

	for x := range intMultipleArray {
		for y := range intMultipleArray[x] {
			fmt.Println(intMultipleArray[x][y])
		}
	}
}
