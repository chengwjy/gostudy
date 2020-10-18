package main

import "fmt"
// import "strings"

// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型
// 这种数据类型叫结构体
// Go语言通过struct实现面向对象
// 使用type 和 struct关键字来实现结构体
type person struct{
	name string
	city string
	age int8
}
type myInt int

func main() {

	// 只有当结构体实例化，才会真正分配内存
	// 也就是实例化后才能使用结构体的字段
	// var 结构体实例 结构体类型
	var p1 person
	p1.name = "迪丽热巴"
	p1.city = "北京"
	p1.age = 18
	fmt.Println(p1)
	fmt.Println(p1.name)

	// 匿名结构体
	// 在实例化的时候定义
	var dog struct{
		name string
		sex bool
	}
	dog.name = "贝贝"
	dog.sex = false
	fmt.Println(dog)

	// 结构体指针
	var p2 = new (person)
	(*p2).name = "tony"
	// 语法糖。结构体类型的指针使用字段与结构体方法一致
	p2.age = 99
	fmt.Println(p2)
	
	// 取结构体的地址进行实例化
	p3 := &person{}
	fmt.Println(p3)

	// 使用键值对初始化
	p4 := person{
		name: "小王子",
	}
	fmt.Println(p4)
	// 使用值的列表初始化
	p5 := &person{
		"小王子",
		"北京",
		33,
	}
	fmt.Println(p5)

	// 实现结构体的构造函数。
	// 因为struct是值类型，如果结构体比较复杂的话，
	// 值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型
	p6 := newPerson("郭靖", "襄阳", 28)
	fmt.Println(p6)

	// 方法
	// go语言种的方法是一种所用与特定类型变量的函数。
	// 这种特定类型变量叫做接收者。
	// func(接受者变量 接收者类型)方法名(参数列表)(返回参数){
	//		函数体
	// }
	// 调用方法
	p6.Dream()
	p6.setAge(99)
	fmt.Println(p6.age)

	var i myInt = 4
	i.printMyint()

}

// 为自定义类型添加方法
// 注意：只能为包内定义的类型添加方法
func (i myInt)printMyint(){
	fmt.Println(i)
}

// 为person类型定义方法
func (p person) Dream(){
	fmt.Println(p.name, "的梦想是学好golang")
}
// 为person类型定义修改年龄的方法
// 指针接收者指的是接收者的类型是指针
func (p *person)setAge(age int8){
	p.age = age
}
func newPerson(name, city string, age int8) *person{
	return &person{
		name: name,
		city: city,
		age: age,
	}
}