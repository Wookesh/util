package util

import "testing"

func TestID_Next(t *testing.T) {
	id := ID{}

	count := 1000000

	data := make([]int32, count)
	for i := 0; i < count; i++ {
		go func(i int) {
			data[i] = id.Next()
		}(i)
	}

	check := make(map[int32]bool)
	for _, elem := range data {
		_, ok := check[elem]
		if ok {
			t.Fail()
		}
		check[elem] = true
	}
}
