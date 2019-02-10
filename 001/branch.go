package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {														// switch后可以没有表达式
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d\n", score)) 		//自动带break，不需要手动写入break
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func main() {
	const filename = "001/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	/**
	 *if中可以赋值后判断， 这里面赋值的变量， 只在if作用域中
	 */
	if con, error := ioutil.ReadFile(filename); err != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", con)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(75),
		grade(82),
		grade(99),
		grade(100),
		grade(101),
		grade(-101),
		)
}
