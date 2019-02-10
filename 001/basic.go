package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3										//go语言没有全局变量， 只有包变量
var ss = "kkk"
var ee = "#EEEEEE"								//包内部变量定义后可以不使用
var bb = true									//这是包内部变量，宝内部变量不能用 := 来赋值

var (
	cc = 4										//可以用()吧所有的变量包含进来
	dd = 7
)


func variableZeroValue() {
	var a int									//不赋初值，系统会自动赋默认值
	var s string								//局部变量定义后， 必须使用，不是会用会编译报错
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4							//可在一行中定义多个变量
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"			//可以不手动指定变量的类型
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"				//用 := 替代 var 来定义变量，这种方式只能在函数内使用
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {									//欧拉公式
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)

	fmt.Printf("%3f\n", cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))	//强制类型转换，类型是需要手动强制转化的，系统不会自动适配类型
	fmt.Println(c)
}

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a * a +  b * b))		//a，b是【没有定义类型】的【常量】，在编译时，会直接替换为需要的类型，所以不需要强制转换
	fmt.Println(filename, c)

	const (
		fiename = "d.txt"
		d, e = 5, 6
	)
	fmt.Println(filename, d , e)
}

func enums() {									//枚举
	const (
		cpp	= 0
		java = 1
		python = 2
		golang = 3
	)
	println(cpp, java, python, golang)

	const (
		c = iota
		js
		erlang
		react
		_
		vue
	)
	println(c, js, erlang, react, vue)
}

func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, cc, dd, ss)

	euler()
	triangle()

	consts()
	enums()
}
