package main

import "fmt"

func heapSort(s []int) {
	length := len(s)
	// 构建大堆顶 i为有叶子节点的节点索引开始
	for i := length/2-1; i >= 0; i-- {
		heapAdjust(s, i, length)
	}
	// 从末尾索引开始，依次与第一个元素交换，并重新构建大堆顶
	for i := length-1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		heapAdjust(s, 0, i)
	}
}

func heapAdjust(s []int, pos int, length int) {
	// 完全二叉树性质，有叶子节点的节点索引为pos，则其左孩子节点索引为2*pos+1 i为索引所以条件为length-1
	for i := 2*pos+1; i <= length-1; i = 2*pos+1 {
		// 当没有右孩子时，如果不加i<length-1，则会报索引错误
		if i < length-1 && s[i] < s[i+1] {
			i++
		}
		// 节点的值大于孩子节点的值，退出循环
		if s[pos] > s[i] {
			break
		}
		// 交换
		s[pos], s[i] = s[i], s[pos]
		// 更新pos节点
		pos = i
	}
}

func main()  {
	s := []int{21, 34, 11, 27, 9, 15, 3, 18}
	heapSort(s)
	fmt.Println(s)
}