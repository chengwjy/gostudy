package main

import "fmt"
import   "strings"
func main() {
	/*
	 * 变量
	 */
	// 变量和常量都不能重新声明
	var name string
	var age int 
	var isOk  bool 
	fmt.Println(name, age, isOk)
	// 批量声明变量
	var (
		a string 
		b int 
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	// 声明变量同时指定初始值
	var book string = "小王子"
	fmt.Println(book)
	// 类型推导，让编译器根据变量的初始值推导出其类型
	var sex, city = 0, "北京"
	fmt.Println(sex, city)
	author := "鲁迅"
	// 短变量声明，只允许用在函数的内部
	fmt.Println(author)
	// 匿名变量，不占用命名空间，不会分配内存
	_ = 123
	
	/*
	 * 常量
	 */
	 // 常量在定义是必须赋值，且不能再次赋值
	const pi = 3.14159
	const (
		system = "windows"
		language = "golang"
	)

	// iota是常量计数器，只能在常量的表达式中使用
	// iota在const关键词出现时被重置为0，const中每
	// 新增一行常量声明将使iota计数一次
	// 使用iota能简化定义，在定义枚举时很有用
	const (
		a1 = iota // 0
		_         // 1
		a2 = 100		  // 插队
		a3 = iota // 3
		a4        // 4
		          // ..
	)
	fmt.Println(a1, a2, a3, a4)

	// 打印d:\go
	fmt.Println("d:\\go")
	demo := `多
行
字
符
串`
	fmt.Println(demo)
	// 打印字符串长度
	fmt.Println(len("你好")) // 6
	// 字符串拼接
	fmt.Println("日" + "本")
	// 字符串分割
	ddd := "how do you do"
	fmt.Println(strings.Split(ddd, " ")) // 得到切片类型

	// 字符
	// byte unit8的别名 AECII码
	// rune int32的别名
	var b1 byte = 'c' 
	fmt.Printf("%T\n", b1) // uint8

	// 条件分支
	score := 65
	if score >= 90{
		fmt.Println("A")
	}else if score >= 80{
		fmt.Println("A")
	}else{
		fmt.Println("C")
	}
	for i:= 0; i < 2; i++{
		fmt.Println(i)
	}
	i := 0
	for i < 2{
		fmt.Println(i)
		i++
	}
	year := 2020
	switch year {
		case 2020:
			fmt.Println("是今年")
			break
		default:
			fmt.Println("不是今年")	
	}
}
