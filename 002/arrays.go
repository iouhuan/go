package main

import "fmt"

/**
1、数组是值类型, 如[3]int 和 [5]int 是不同的类型， 参数是传不进去的
2、go中一般不使用数组
 */

// 打印， 数组是值类型， 所以传入的必须是[5]int, 类似[3]int是不行的
func printArray(arr [5]int){
	arr[0] = 100						// 因为go是值传递， 这地方改变内部， 改变不了main中的数组
	for i, v := range arr{
		fmt.Println(i, v)
	}

}

func printArray2(arr *[5]int) {
	arr[0] = 100						// 因为go是引用传递, 会改变main中的数组
	for i, v := range arr{
		fmt.Println(i, v)
	}
}

func printArray3(arr []int) {			// 利用slice实现， 会改变main中的数组
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}


func main() {
	var arr1 [5]int						// 默认为0
	arr2 := [3]int{1, 3, 5}				// 利用:=定义， 要初始化
	arr3 := [...]int{2, 4, 6, 8, 10}	// 不定义数组长度

	var grid [4][5]int					// 二维数组， 默认为0, 4行5列

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {	//常规循环
		fmt.Println(arr3[i])
	}

	for j := range arr3 {				// 只有key
		fmt.Println(arr3[j])
	}

	for k, v := range arr3 {			// 有key也有value
		fmt.Println(k, v)
	}

	for _, value := range arr3 {		// 只要值，不要key
		fmt.Println(value)
	}

	printArray(arr1)
	//printArray(arr2)					// 这地方因为是[3]int， 是传不进去的
	printArray(arr3)					// 长度自动计算出来时5， 是可以传的

	fmt.Println(arr1, arr3)				// printArray内部改造了数组索引为0的值， 但不影响arr1和arr3

	fmt.Printf("-------------\n")
	printArray2(&arr1)
	printArray2(&arr3)

	fmt.Println(arr1, arr3)				// printArray2内部改造了数组索引为0的值， 会影响arr1和arr3

	fmt.Printf("-------------\n")
	printArray3(arr1[:])
	printArray3(arr3[:])

	fmt.Println(arr1, arr3)				// printArray3是切片实现， 会改变arr1和arr3

}
