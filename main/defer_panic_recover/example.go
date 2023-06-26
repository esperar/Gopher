package defer_panic_recover

import (
	"fmt"
	"os"
)

/**
defer: 함수를 바로 실행하지 않고 종료 시점에 실행
panic: 실행 시점에서 현재 함수의 defer() 함수를 실행 후 리턴
recover: panic()에 의한 에러를 복구
*/

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		// error가 발생한 시점에 종료
		panic(err)
	}

	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))

	openFile("Invalid.txt")

}

func openFile(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
