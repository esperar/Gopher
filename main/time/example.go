package time

import (
	"fmt"
	"time"
)

func timeParsing() {
	timeStr := "2023-04-04 23:01:29"
	parseTime, err := time.Parse("2006-01-02 15:04:05", timeStr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parseTime)
}
