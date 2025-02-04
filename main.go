package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		time.Sleep(3 * time.Second)
		fmt.Printf("The value of data is: %v\n", data)
	}
}
