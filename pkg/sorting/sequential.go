package sorting

import (
	"sort"
)

func SequentialSort(arr [][]int) {
	for _, arrEle := range arr {
		sort.Ints(arrEle)
	}
}
