package main

import (
	"log"
	"os"
)

// go는 error 인터페이스를 가지고 있다.

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C:\\temp\\1.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	println(f.Name())

	// 사용자 정의 에러를 잡을 때
	switch err.(type) {
	default: // no error
		println("ok")
	//case MyError:
	//	log.Print("Log my error")
	case error:
		log.Fatal(err.Error())
	}
}
