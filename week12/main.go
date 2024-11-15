package main

import (
	"fmt"
	"time"
)

func main() {
	var dates = [3]time.Time{time.Unix(1, 0), time.Unix(2, 0)}
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println(dates)
	//fmt.Printf("%#v\n", dates)

	for i := 0; i < len(dates); i++ {
		fmt.Println(i, dates[i])
	}
}
