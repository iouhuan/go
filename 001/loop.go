package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
1、for if后的条件没有括号
2、if条件里也可以定义变量
3、没有while
4、switch不需要break，也可以直接switch多个条件
 */

// 整数转为二进制，存的顺序是逆序的， 如13是1011， 存储时1101
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

// 读文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {		// 没有初始条件， 也没有递增条件， 只有结束条件（两个分号都可以不用写）， 相当于其他语言中的while， go语言中没有while
		fmt.Println(scanner.Text())
	}
}

func forever(){
	for {						// 三种条件都没有也是可以的， 就是死循环了
		fmt.Println("abcd")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),		//101
		convertToBin(13),	//1101
	)// 如果小括号最后是独立的行 那么上一行必须有逗号， 如果不是， 可以不用逗号

	printFile("001/abc.txt")

	forever()
}
