package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
1、函数可以返回多个值, 返回值类型写在最后面
2、返回多个值时，能起名字【仅用于非常简单的函数】，对于调用者而言，没有区别
3、函数可以作为参数
4、没有其他语言中花哨的默认参数和可选参数，都是非常严格的
5、可变参数列表
6、go中指针不能计算， 概念和c中相同, 写法上是：变量名 *类型， 如 var a *int
7、参数的传递是值传递， copy一份
 */
func eval(a, b int, op string) (int, error) {		// error是一种类型
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div2(a, b)							// 用不到的变量， 可以用_代替， go中只要定义就必须使用
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s\n", op)
	}
}

func div(a, b int) (int, int){
	return a / b, a % b
}

//可以定义返回值起名字， 在函数体内对变量进行计算， 最后只return即可，不用return值
//简单的函数这样比较方便， 如果结构体的代码比较多， 不建议这样， 不好找到那句是赋值的
func div2(a, b int) (q, r int){
	q = a / b
	r = a % b
	return
}

// go中函数是一等公民， 参数也可以是函数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

// main.pow
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可变参数列表
func sum(numbers ...int) int{
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

// 利用指针交换两个变量的值
func swap(a, b *int){
	*b, *a = *a, *b
}

// 常规函数交换实现变量值交换
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(
		eval(3, 4, "*"),
		)
	fmt.Println(
		div(13, 3),
		)

	q, r := div2(13, 3)
	fmt.Println(q, r)

	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		fmt.Println(result);
	}

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		// 匿名函数, main.main.func1
		func(a int , b int) int{
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 3, 4
	c, d = swap2(c, d)
	fmt.Println(c, d)
}
