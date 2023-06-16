package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var 선언
	var i1 int = 10
	var s1 string = "Hello"

	fmt.Println(i1)
	fmt.Println(s1)

	// 타입 생략 가능
	var i2 = 10
	var s2 = "Hello"

	fmt.Println(i2)
	fmt.Println(s2)

	// := 를 이용한 변수 선언
	i3 := 10
	s3 := "Hello"

	fmt.Println(i3)
	fmt.Println(s3)

	// 다수의 변수를 동시에 선언합니다.
	var j1, j2, j3 = 1, 2, 3
	s4, s5, s6 := "string1", "string2", "string3"

	fmt.Println(j1, j2, j3)
	fmt.Println(s4, s5, s6)

	// var ()를 이용한 변수 선언
	var (
		i7            = 10
		i8            = 20
		i9            = 30
		s10, s11, s12 = "string1", "string2", "string3"
	)

	fmt.Println(i7, i8, i9)
	fmt.Println(s10, s11, s12)

	const i = 1
	const s string = "asdfg"

	// iota가 0이 할당되고 순서대로 증가
	const (
		A = iota
		B
		C
		D
	)

	fmt.Println(i, s)
	fmt.Println(A, B, C, D)
}
