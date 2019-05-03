package main

import "fmt"

func main() {
	//初始化切片
	var s = heap{
		5, 8, 12, 3, 7, 6, 2, 1, 9,10, 14,
	}

	fmt.Println("====排序前====")
	fmt.Println(s)
	Heapsort(s)
	fmt.Println("====排序后====")
	fmt.Println(s)
}

type heap []int

type Heap interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
	Sink(i, n int)
}

func Heapsort(s heap) {
	N := s.Len()
	//构造堆
	s.MakeHeap(N-1)

	//堆排序
	//下沉排序
	for N > 0 {
		//将根节点下沉到末尾
		s.Swap(0, N-1)
		N--
		s.Sink(0, N-1)
	}
}

func (h heap) Sink(i, n int) {
	for 2*(i+1) <= n {
		j := 2 * i + 1
		if j < n && h.Less(j, j+1) {
			j++
		}
		if !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h heap) MakeHeap(n int) {
	for k := n / 2; k >= 0; k-- {
		h.Sink(k, n)
	}
}
func (h heap) Len() int {
	return len(h)
}

func (h heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heap) Less(i, j int) bool {
	return h[i] < h[j]
}
