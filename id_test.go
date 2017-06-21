package util

import (
	"sync"
	"testing"
)

func TestID_Next(t *testing.T) {
	id := ID{}

	count := 1000000

	data := make([]int32, count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			data[i] = id.Next()
		}(i)
	}

	wg.Wait()

	check := make(map[int32]bool)
	for _, elem := range data {
		_, ok := check[elem]
		if ok {
			t.Fail()
		}
		check[elem] = true
	}
}
