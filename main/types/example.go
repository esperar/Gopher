package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strInt := "100"
	intStr := string(strInt)
	fmt.Println(intStr, reflect.TypeOf(intStr)) // 100 string

	i, err := strconv.Atoi(strInt)
	fmt.Println(i, err, reflect.TypeOf(i)) // 100 <nil> int

	strInt = "987654321"
	i8, err := strconv.ParseInt(strInt, 0, 8)
	i16, err := strconv.ParseInt(strInt, 0, 16)
	i32, err := strconv.ParseInt(strInt, 0, 32)
	i64, err := strconv.ParseInt(strInt, 16, 64)

	fmt.Println(i8, err, reflect.TypeOf(i8))   // 127 <nil> int64
	fmt.Println(i16, err, reflect.TypeOf(i16)) // 32767 <nil> int64
	fmt.Println(i32, err, reflect.TypeOf(i32)) // 987654321 <nil> int64
	fmt.Println(i64, err, reflect.TypeOf(i64)) // 40926266145 <nil> int64
}
