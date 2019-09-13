package main

import "fmt"

func merge(arr []int, low int, mid int, high int) {  // low: 0 mid: 3 high: 7
	// 长度是索引值+1
	temp := make([]int, high-low+1)
	// 把无序的arr赋给temp
	for i := low; i <= high; i++ {
		temp[i-low] = arr[i]  // temp: [1, 3, 5, 7, 2, 4, 6, 8]
	}

	left := low   // 0
	right := mid + 1  // 4

	for i := low; i <= high; i++ {   // i是arr的索引
		// 处理最后一个数(有可能是temp的最后一个或中间的那个)存放到arr最后一个位置
		if left > mid {
			arr[i] = temp[right-low]
			break
		} else if right > high {
			arr[i] = temp[left-low]
			break
		// 处理左右数组中每个值的顺序
		} else if temp[left-low] > temp[right-low] {
			arr[i] = temp[right-low]
			right++
		} else if temp[left-low] < temp[right-low] {
			arr[i] = temp[left-low]
			left++
		}
	}
}

func mergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	// 递归向下
	mid := (high + low) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	// 归并向上
	merge(arr, low, mid, high)
}

func main() {
	arr1 := []int{3, 5, 1, 7}
	arr2 := []int{4, 8, 2, 6}
	arr := make([]int, len(arr1)+len(arr2))
	copy(arr[0:len(arr1)], arr1)
	copy(arr[len(arr1):], arr2)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	/*s := "abcd"
	fmt.Println(letterPermutation(s))*/
}

func letterPermutation(s string) []string {
	str := make([]string, 0)
	str = helper([]byte(s), 0, str)
	return str
}

func helper(bt []byte, start int, str []string) []string {
	if start == len(bt) {
		str = append(str, string(bt))
	} else {
		for i := start; i < len(bt); i++ {
			if i != start {
				bt[start], bt[i] = bt[i], bt[start]
			}
			str = helper(bt, start+1, str)
			if i != start {
				bt[start], bt[i] = bt[i], bt[start]
			}
		}
	}
	return str
}
