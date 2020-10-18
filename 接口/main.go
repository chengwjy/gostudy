package main
import "fmt"
type dog struct{}
func (d dog)say(){
	fmt.Println("汪汪汪")
}
type cat struct{}
// 值接收者实现接口
func (d cat)say(){
	fmt.Println("喵喵喵")
}

type person struct{}
// 指针接收者实现接口
func (p *person)say(){
	fmt.Println("啊啊啊")
}
// 接口不管你是什么类型，它只管你要实现什么方法
// 定义一个抽象的类型。只要实现了say()的类型都可以成为sayer
type sayer interface{
	say()
}
// 打的函数
func da(arg sayer){
	arg.say()
}
func main(){
	// 在go语言中，接口是一种类型，是一种抽象的类型
	// 接口做的事情就像是定义一个协议(规则)
	// 只要一台辑器有洗衣服和帅干的功能
	// 我们就称它为洗衣机。不关心属性(数据),只关心行为(方法)
	//  一个对象只要全部实现了接口中的方法，那么就实现了该接口
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)

	// 使用值类型接收者实现接口
	// 类型的值和类型的指针都能够保存到接口
	var s sayer
	c2 := cat{}
	s = &c2
	s = c2
	fmt.Println(s)

	// 使用指针接收者实现接口。
	// 只有指针类型能够保存到接口变量中
	p1 := &person{}
	s = p1
	fmt.Println(s)

	// 空接口
	// 空接口是指没有定义任何方法的接口
	// 因此任何类型都实现了空接口
	// 空接口的类型可以存储任意类型的变量
	var x interface{}
	x = 100
	x = true
	x = "你猜我是谁"
	fmt.Println(x)
	// 类型断言
	_, ok := x.(string)
	
	fmt.Println(ok) // 如果x是string类型，ok是true，否则是flase

	// 使用空接口作为函数的参数
	myPrintln("ddd")	
	// 空接口作为map的value
	var m = make(map[string]interface{})
	m["name"] = "娜扎"
	m["age"] = 18
	myPrintln(m)
}


func  myPrintln(age interface{}){
	fmt.Println(age)
}