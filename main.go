package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	fmt.Printf("time1 %v\n", time1.Format(time.RFC3339))
	var time2 = time.Now().UTC()
	fmt.Printf("time2 %v\n", time2)
	fmt.Printf("time2 %v\n", time2.Format(time.RFC3339))
	var time3 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time3 %v\n", time3)
	fmt.Printf("time3 %v\n", time3.Format(time.RFC3339))
	var time4 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.Local)
	fmt.Println(time.Local, time.UTC)
	fmt.Printf("time4 %v\n", time4)
	fmt.Printf("time4 %v\n", time4.Format(time.RFC3339))
}
