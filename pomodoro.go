package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting Pomodolo Technique")

	if len(os.Args) < 2 {
		fmt.Println("Please input the number of set.")
		return
	}
	set, err := strconv.Atoi(os.Args[1]) //set数

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(time.Now())
	fmt.Println(set)

	const duration_minute int64 = 25
	const interval_minute int64 = 5

	ticker := time.NewTicker(time.Duration(duration_minute) * time.Minute)
	for i := 0; i < set; i++ {
		t := <-ticker.C
		fmt.Println(t)
		fmt.Println("Woriking time is done. 5m interval is starting.")
		interval_ticker := time.NewTicker(time.Duration(interval_minute) * time.Minute)
		t = <-interval_ticker.C
		if i < set-1 {
			fmt.Println(t)
			fmt.Println("5m interval is done. Next working time is starting")
		}
		interval_ticker.Stop()
	}
	ticker.Stop()
	fmt.Println("Pomodolo Technique is done.")
}
