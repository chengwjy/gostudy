package main

import "fmt"
// import   "strings"

// 函数接收可变参数
// 当多个参数是相同类型时，可以只将数据类型写在最后一个参数里面
func intSum(nums ...int)(sum, len int){
	// 可变参数在函数体中是切片类型
	for _, num := range(nums){
		// 当返回值有变量名时，不需要重新声明
		sum += num
	}
	len = cap(nums)
	// 函数在即将要结束的时候按照逆序的方式，执行被defer修饰的语句，通常用来处理资源的释放问题
	defer fmt.Println(999)
	defer fmt.Println(998)
	return sum, len
}
func main() {
	// go语言中支持函数、匿名函数和闭包，并且函数
	// 在go语言中属于“一等公民”	
	sum, len := intSum(1,2)
	fmt.Println(sum, len)

	// 函数是可以作为变量的
	testFunc := intSum
	fmt.Printf("%T\n", testFunc)


	// 匿名函数
	// 匿名函数就是没有函数名的函数，匿名函数多用于实现回调函数和闭包
	func(x int){
		fmt.Println(x)
	}(10)

	// 闭包
	// 闭包指的是一个函数与其相关的引用环境组合而成的实体
	// 简单来说，闭包=函数+引用环境
	// funca 此时就是闭包
	// 判断函数是不是闭包，就要判断函数内部有没有引用外层的作用域的变量
	funca := a()
	funca()

	b()
	// recover()必须搭配defer使用
	// defer一定要在可能引发painc的语句之前定义
}

// 把函数作为返回值
func a() func(){
	name := "古力娜扎"
	return func(){
		fmt.Println(name)
	}
}

func b(){
	// defer注册
	defer func(){
		// 收集相关错误信息
		err := recover()
		if err != nil{
			fmt.Println("func b err")
		}		
	}()

	// 触发painc
	panic("painc in b")
}