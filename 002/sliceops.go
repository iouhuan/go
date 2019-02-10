package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("val=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	// 第一种定义方式
	var s []int							// Zero value for slice is nil
	println(s == nil)

	for i := 0; i < 100; i++ {
		printSlice(s)					// go会自己扩充cap，已有的cap*2
		s = append(s, 2 * i + 1)
	}

	fmt.Println(s)

	// 第二种定义方式
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// 第三种定义方式
	s2 := make([]int, 16)
	printSlice(s2)

	// 第四种定义方式， 32是cap
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Coping slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	printSlice(s2[:3]);
	printSlice(s2[4:]);
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
