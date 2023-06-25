package main

import "fmt"

func main() {

	// 맵 선언
	intMap := map[int]string{}
	intMap[100] = "A"
	intMap[101] = "B"
	fmt.Printf("%+v\n", intMap)
	fmt.Println(intMap[100])

	// make로 맵 선언
	var stringMap = make(map[string]string)
	stringMap["A"] = "Hello"
	stringMap["B"] = "World"
	fmt.Printf("%+v\n", intMap) // map[100:A 101:B]
	fmt.Println(stringMap["B"]) // World

	// 선언과 동시에 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSGT": "Microsoft",
		"FB":   "FaceBook",
	}

	fmt.Printf("%+v\n", tickers)
	fmt.Println(tickers["FB"])

	// 맵에 데이터가 존재 하는지 확인
	val, exists := tickers["FB"]
	fmt.Println(val, exists) // FaceBook true

	val, exists = tickers["NOT_EXISITS"]
	fmt.Println(val, exists) // false

	// range 문으로 key, value 반복 처리
	for key, value := range tickers {
		fmt.Println(key, value)
	}
}
