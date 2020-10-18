package main

import "fmt"
// import "strings"
func main() {
	// 指针
	// 区别于C/C++中的指针，Go语言中的
	// 指针不能进行偏移和运算，是安全指针
	// Go语言中的函数传参都是值拷贝
	// 当我们需要修改某个变量，可以创建一个指向
	// 该变量地址的指针变量

	// Go语言中的值类型（int，float，bool，string，array，struct）
	// 都有对应的指针类型
	a := 10
	b := &a
	fmt.Println(b) //0xc000012090
	fmt.Println(*b) //10

	// new和make的区别
	// new是一个内置函数，它的函数签名如下
	// func new(Type) *Type
	// new函数不太常用，使用new函数得到的是一个类型的指针
	// 并且该指针对应的值为该类型的零值
	c := new(int)
	fmt.Println(*c) //0
	// make也用于内存分配，区别于new
	// 它只用于sloce、map以及chan的内存创建
	// 而且它返回的类型就是折三个类型本身，而不是它们的指针类型
	// 因为这三种类型就是引用类型，所以没必要返回它们的指针了
	// 它的函数签名如下
	// func make(t Type, size ...IntegerType) Type
	// make函数无可替代，我们使用sloce、map以及chan 的时候
	// 都需要用make进行初始化，然后才可以对它进行操作
	x := make([]int, 4)
	fmt.Println(x)

	// 自定义类型
	// go语言中可以通过type关键字来自定义类型
	// 自定义类型定义了一个全新的类型，我们可以基于内置的基本类型定义
	// 也可以通过struct定义
	// 通过type关键词的定义，MyInt就是一中新的类型，它具有int的特性
	type MyInt int
	var y MyInt = 55
	fmt.Println(y)

	// 类型别名
	type TheInt = int
	var z TheInt = 55
	fmt.Println(z)
	// 我们之前见过的type就是类型别名
	// 它们的定义是 type byte = uint8
} 