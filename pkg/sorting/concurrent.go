package sorting

import (
	"sort"
	"sync"
)

func ConcurrentSort(arr [][]int) {
	var wg sync.WaitGroup

	for _, arrEle := range arr {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()
			sort.Ints(arr)
		}(arrEle)
	}

	wg.Wait()
}
