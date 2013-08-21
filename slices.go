package main

import (
	"fmt"
	"time"
)

func main() {

	const sz = 100000
	sl1 := make([]int32, sz)
	sl2 := make([]int32, sz)

	for j := 0; j < sz; j++ {
		sl1[j] = int32(j)
		sl2[j] = int32(j * 2)
	}

	start := time.Now()
	var num int64 = 0
	for _, i := range sl1 {
		for _, j := range sl2 {
			num += int64(i * j)
		}
	}

	et := time.Since(start)
	fmt.Println(et)
	fmt.Println(num)
}

// EOF
