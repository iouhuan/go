package main

import "fmt"

/**
1、切片slice是半闭半开
2、slice本身是没有数据，是对底层array的一个view(视图), 哪怕是多次reslice， 都是原数组的一个视图
3、slice 可以向后扩展， 不能向前扩展，s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
 */


func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[2:6]
	fmt.Println("arr[2:6] = ", s)

	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Printf("After updateSlice s1 is: \n")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Printf("After updateSlice s2 is: \n")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Printf("Reslice: \n")			// 多次reslice， 都是原数组的一个视图
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	s2 = s2[1:]
	fmt.Println(s2)
	updateSlice(s2)
	fmt.Println(s2)

	fmt.Println("arr = ", arr)


	// slice 可以向后扩展(向后扩展[:n]， cap是剩余的所有坑位)， 不能向前扩展([n:])
	// s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	// cap方法： Slice: the maximum length the slice can reach when resliced;
	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr = ", arr)
	s3 := arr[2:6]
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	s4 := s3[3:5]
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))

	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)

	fmt.Println("s5, s6, s7 = ", s5, s6, s7)
	// s7 已经不再是arr的view， 是系统重新分配的一个array
	// 添加元素时如果超过了cap(s), 系统会重新分配更大的底层数组， 并且会把原来的直接copy一份过去
	// 由于值传递的关系， 会改变slice的ptr、len、cap, 必须接受append的返回值
	fmt.Println("arr = ", arr)
}
