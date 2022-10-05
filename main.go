// Entry point for a sample
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	res := make([]int, 0)

	for i := 0; i < 3; i++ {
		go func(num int) {
			ch <- multiply(num)
		}(i)
	}

	res = append(res, <-ch)
	fmt.Println(res)
}

func multiply(num int) int {
	time.Sleep(2 * time.Second)
	return num * num
}
