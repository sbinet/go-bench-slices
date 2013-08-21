package main

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) {

	sl1 := make([]int32, sz)
	sl2 := make([]int32, sz)

	for j := 0; j < sz; j++ {
		sl1[j] = int32(j)
		sl2[j] = int32(j * 2)
	}

	b.ResetTimer()

	var num int64 = 0
	for _, i := range sl1 {
		for _, j := range sl2 {
			num += int64(i * j)
		}
	}

	if num != expected {
		b.Errorf("expected [%v] got [%v]", expected, num)
	}
}

func BenchmarkArray(b *testing.B) {

	sl1 := [sz]int32{}
	sl2 := [sz]int32{}

	for j := 0; j < sz; j++ {
		sl1[j] = int32(j)
		sl2[j] = int32(j * 2)
	}

	b.ResetTimer()

	var num int64 = 0
	for _, i := range sl1 {
		for _, j := range sl2 {
			num += int64(i * j)
		}
	}

	if num != expected {
		b.Errorf("expected [%v] got [%v]", expected, num)
	}
}

// EOF
