package main

import "fmt"

/**
1、go语言中， empty和nil是可以参与运算的
2、获取元素m[key]
3、key不存在时，获得Value类型的初始值Zero value
4、用value, ok := m[key]来判断是否存在key
5、delete删除一个key
6、使用range遍历key，或者key，value对
7、不保证遍历的顺序， 需利用slice手动进行排序
8、使用len获得map的元素个数
9、map的key
9.1 map使用哈希表，必须可以比较相等
9.2 除了slice、map、function的内建类型都可以作为key
9.3 struct类型不包含上述字段，也可以作为key
 */
func main() {
	// 第一种定义方式
	m := map[string]string {
		"name": "ccmouse",
		"course": "golang",
		"site": "imooc",
		"quality": "notbad",
	}

	fmt.Println(m);

	// 第二种定义方式
	m2 := make(map[string]int)					// m2 == empty map
	fmt.Println(m2)

	// 第三种定义方式
	var m3 map[string]int						// m3 == nil
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {						// key在map中是无序的， 每次打印出顺序可能会不同
		fmt.Println(k, v)
	}

	fmt.Println("Getting value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {		// 返回的第二个值是是否存在该key
		fmt.Println(causeName)					// 没有该key会返回Value类型的Zero value : ""
	} else {
		fmt.Println("key don't exit.")
	}

	fmt.Println("Deleting value")
	fmt.Println(m);

	delete(m, "name")
	fmt.Println(m)

	fmt.Println(m3["key"])						// 没有该key会返回Value类型的Zero value : 0

}
