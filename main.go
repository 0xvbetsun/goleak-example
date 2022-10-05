// Entry point for a sample
package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	cnt := 3
	res := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			res[num] = multiply(num)
		}(i)
	}
	wg.Wait()
	out, _ := json.Marshal(res)
	fmt.Println(string(out))
}

func multiply(num int) int {
	time.Sleep(2 * time.Second)
	return num * num
}
