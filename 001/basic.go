package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Printf("%d %q\n", a, s)
}

func main(){
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
}