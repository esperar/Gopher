package main

import (
	"fmt"
	"strings"
)

func main() {

	// 문자열 합치기 Join
	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, ":")) // a:b:c

	// 문자열 대치 (replace)
	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1)
	fmt.Println(r) // a_b_c
}
