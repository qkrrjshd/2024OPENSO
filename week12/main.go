package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time
	//dates[1] = time.Unix(1447920000, 0)
	//fmt.Println(dates)
	var dates = [3]time.Time{time.Unix(1, 0), time.Unix(2, 0)}
	fmt.Println(dates[0], dates[1], dates[2])
}
