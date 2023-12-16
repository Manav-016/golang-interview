package api_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/golang-interview/pkg/sorting"
)

func generateRandomLargeSubarrays() [][]int {
	var largeArray [][]int
	for i := 0; i < rand.Intn(100)+10; i++ {
		subArraySize := rand.Intn(100) + 1
		subArray := make([]int, subArraySize)
		for j := range subArray {
			subArray[j] = rand.Intn(10000)
		}
		largeArray = append(largeArray, subArray)
	}
	return largeArray
}

func sortRandomLargeSubarrays(arr [][]int) [][]int {
	for _, subArr := range arr {
		sort.Ints(subArr)
	}
	return arr
}

func TestSequentialSort(t *testing.T) {

	randomLargeArray := generateRandomLargeSubarrays()

	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{}},
			expected: [][]int{{}},
		},
		{
			input:    [][]int{{1}, {2}, {3}},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			input:    [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			input:    [][]int{{-1, -3, -2}, {0, -1, 1}},
			expected: [][]int{{-3, -2, -1}, {-1, 0, 1}},
		},
		{
			input:    [][]int{{5, 5, 5}, {2, 1, 2}, {3, 3, 1}},
			expected: [][]int{{5, 5, 5}, {1, 2, 2}, {1, 3, 3}},
		},
		{
			input:    [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}},
			expected: [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}},
		},
		{
			input:    randomLargeArray,
			expected: sortRandomLargeSubarrays(randomLargeArray),
		},
	}

	for _, tc := range testCases {
		sorting.SequentialSort(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("SequentialSort failed. Expected %v, got %v", tc.expected, tc.input)
		}
	}
}

func TestConcurrentSort(t *testing.T) {

	randomLargeArray := generateRandomLargeSubarrays()

	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{}},
			expected: [][]int{{}},
		},
		{
			input:    [][]int{{1}, {2}, {3}},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			input:    [][]int{{3, 2, 1}, {6, 5, 4}, {9, 8, 7}},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			input:    [][]int{{-1, -3, -2}, {0, -1, 1}},
			expected: [][]int{{-3, -2, -1}, {-1, 0, 1}},
		},
		{
			input:    [][]int{{5, 5, 5}, {2, 1, 2}, {3, 3, 1}},
			expected: [][]int{{5, 5, 5}, {1, 2, 2}, {1, 3, 3}},
		},
		{
			input:    [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}},
			expected: [][]int{{1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}},
		},
		{
			input:    randomLargeArray,
			expected: sortRandomLargeSubarrays(randomLargeArray),
		},
	}

	for _, tc := range testCases {
		sorting.ConcurrentSort(tc.input)
		if !reflect.DeepEqual(tc.input, tc.expected) {
			t.Errorf("Concurrent failed. Expected %v, got %v", tc.expected, tc.input)
		}
	}
}
