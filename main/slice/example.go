package main

import "fmt"

func main() {

	// 슬라이스 생성
	letter := []string{"a", "b", "c"}
	fmt.Println(letter) // [a, b, c]

	// make를 이용한 slice 생성
	s1 := make([]int, 0)
	fmt.Println(s1, len(s1), cap(s1)) // [], 0, 0

	s2 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s2)) // [], 0, 5

	s3 := make([]int, 3, 5)
	fmt.Println(s2, len(s2), cap(s3)) // [0, 0, 0, 0, 0], 5, 5

	// append는 len으로 설정한 값 뒤에 추가 됨
	intArray := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 100)
	s2 = append(s2, intArray...) // slice 끼리 더할때는 ... 추가
	s3 = append(s3, 100, 101, 102)

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	// copy를 위해서 생성한 슬라이스는 기존 슬라이스와 크기가 동일해야 함
	letterCopy := make([]string, 0)
	copy(letterCopy, letter)
	fmt.Println(letterCopy) // []

	letterCopy2 := make([]string, len(letter), len(letter))
	copy(letterCopy2, letter)
	fmt.Println(letterCopy2) // [a, b, c]

	letterCopy2[2] = "="
	fmt.Println(letter)      // [a, b, c]
	fmt.Println(letterCopy2) // [a, b, =]

	// 삭제, 추출
	integers := []int{1, 2, 3, 4, 5}
	sub1 := integers[1:4]
	sub2 := integers[2:4]
	fmt.Println(integers, len(integers), cap(integers)) // [1,2,3,4,5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1))             // [2,3,4] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2))             // [3,4] 2 3

	sub1[2] = 100
	fmt.Println(integers, len(integers), cap(integers)) // [1 2 3 100 5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1))             // [2 3 100] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2))             // [3 100] 2 3
}
