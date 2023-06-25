package structs

import "fmt"

// go는 객체지향언어가 가지는 클래스 객체 상속 개념이 없다.
// 일반적인 클래스가 필드랑 메서드를 함께 가지는 것과 달리 Go의 구조체는 필드만 가진다.
type person struct {
	name string
	age  int
}

func main() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Kim"
	p.age = 10

	fmt.Println(p)

	// 초기화
	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}

	// new 를 이용한 초기화

	p3 := new(person)
	p3.name = "Lee" // p가 포인터라도 .을 사용한다.

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
