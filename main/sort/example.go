package sort

import (
	"fmt"
	"sort"
)

type myDataType struct {
	name string
	age  int
}

func main() {
	mySlice := make([]myDataType, 0)
	mySlice = append(mySlice, myDataType{"김희망", 18})
	mySlice = append(mySlice, myDataType{"김망희", 19})
	mySlice = append(mySlice, myDataType{"김김김", 20})

	fmt.Println(mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].age < mySlice[j].age
	})

	fmt.Println(mySlice)
}
