package main

import "fmt"

func shellSort(s []int) {
	increment, length := 1, len(s)
	for increment < length / 3 {
		increment = 3 * increment + 1
	}
	for increment >= 1 {
		for i := increment; i < length; i++ {
			for j := i; j >= increment && s[j] < s[j-increment]; j -= increment {
				s[j], s[j-increment] = s[j-increment], s[j]
			}
		}
		increment /= 3
	}
}

func quickSort(s []int) {
	qSort(s, 0, len(s)-1)
}

func qSort(s []int, low, high int) {
	if low < high {
		pivot := partition(s, low, high)
		qSort(s, pivot+1, high)
		qSort(s, low, pivot-1)
	}
}

func partition(s []int, low, high int) int {
	// 三数取中，保证pivotKey是数列中的中间值
	mid := low + (high - low) / 2
	if s[low] > s[high] {
		s[low], s[high] = s[high], s[low]
	}
	if s[mid] > s[high] {
		s[mid], s[high] = s[high], s[mid]
	}
	if s[mid] > s[low] {
		s[mid], s[low] = s[low], s[mid]
	}
	pivotKey := s[low]
	for low < high {
		for low < high && s[high] >= pivotKey {
			high -= 1
		}
		// 当右侧值小于pivotKey时，左侧的low索引替换成high索引的小值
		s[low] = s[high]
		for low < high && s[low] <= pivotKey {
			low += 1
		}
		// 当左侧值大于pivotKey时，右侧的high索引替换成low索引的大值
		s[high] = s[low]
	}
	s[low] = pivotKey
	return low
}

func main() {
	s := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	shellSort(s)
	//quickSort(s)
	fmt.Println(s)
}
