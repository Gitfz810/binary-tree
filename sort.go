package main

import "fmt"

func shellSort(s []int) {
	increment, length := 1, len(s)
	for increment < length / 3 {
		increment = increment * 3 + 1
	}

	for increment >= 1 {
		for i := increment; i < length; i++ {
			for j := i; j >= increment; j -= increment {
				if s[j] < s[j-increment] {
					s[j], s[j-increment] = s[j-increment], s[j]
				}
			}
		}
		increment /= 3
	}
}

func heapSort(s []int) {
	length := len(s)
	for i := length/2-1; i >= 0; i-- {
		heapAdjust(s, i, length)
	}

	for i := length-1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		heapAdjust(s, 0, i)
	}
}

func heapAdjust(s []int, pos int, length int) {
	for i := 2*pos+1; i <= length-1; i = 2*pos+1 {
		if i < length-1 && s[i] < s[i+1] {
			i++
		}

		if s[pos] > s[i] {
			break
		}

		s[pos], s[i] = s[i], s[pos]
		pos = i
	}
}

func mergeSort(s []int, low, high int) {
	if low >= high {
		return
	}

	mid := (high + low) / 2
	mergeSort(s, low, mid)
	mergeSort(s, mid+1, high)
	merge(s, low, mid, high)
}

func merge(s []int, low int, mid int, high int) {
	temp := make([]int, high-low+1)
	for i := low; i <= high; i++ {
		temp[i-low] = s[i]
	}

	left := low
	right := mid + 1

	for i := low; i <= high; i++ {
		if right > high {
			s[i] = temp[left-low]
			left++
		} else if left > mid {
			s[i] = temp[right-low]
			right++
		} else if temp[left-low] < temp[right-low] {
			s[i] = temp[left-low]
			left++
		} else {
			s[i] = temp[right-low]
			right++
		}
	}
}

func quickSort(s []int) {
	qSort(s, 0, len(s)-1)
}

func qSort(s []int, low int, high int) {
	if low < high {
		pivot := partition(s, low, high)
		qSort(s, pivot+1, high)
		qSort(s, low, pivot-1)
	}
}

func partition(s []int, low int, high int) int {
	mid := low + (high-low) / 2
	if s[mid] > s[high] {
		s[mid], s[high] = s[high], s[mid]
	}
	if s[low] > s[high] {
		s[low], s[high] = s[high], s[low]
	}
	if s[mid] > s[low] {
		s[mid], s[low] = s[low], s[mid]
	}
	pivotKey := s[low]

	for low < high {
		for low < high && s[high] > pivotKey {
			high--
		}
		s[low] = s[high]
		for low < high && s[low] < pivotKey {
			low++
		}
		s[high] = s[low]
	}
	s[low] = pivotKey
	return low
}

func main()  {
	s := []int{19, 27, 6, 32, 12, 22, 26, 17, 23}
	fmt.Println(s)
	//shellSort(s)
	//heapSort(s)
	//mergeSort(s, 0, len(s)-1)
	quickSort(s)
	fmt.Println(s)
}
