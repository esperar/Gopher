package main

import (
	"fmt"
	"strings"
)

func changeValue(message *string) {
	*message = fmt.Sprintf("your name is %s", *message)
}

func addTen(value int) int {
	return value + 10
}

func sum(values ...int) int {
	result := 0
	for _, value := range values {
		result += value
	}

	return result
}

func divideStrings(value string) (string, string) {
	values := strings.Split(value, "_")
	return values[0], values[1]
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	// 함수 호출
	fmt.Println(addTen(100))

	// slice 변수 호출
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 다중 변수 반환
	prefix, suffix := divideStrings("Your_Name")
	fmt.Println(prefix, suffix)

	// 포인터 전달
	message := "Esperer"
	changeValue(&message)
	fmt.Println(message)

	// 익명 함수
	sum2 := func(x int, y int) int {
		return x + y
	}

	fmt.Println(sum2(5, 10))

	// 일급 함수
	add := func(i int, j int) int {
		return i + j
	}

	minus := func(i int, j int) int {
		return i - j
	}

	fmt.Println(calc(add, 5, 10))
	fmt.Println(calc(minus, 10, 5))
}
