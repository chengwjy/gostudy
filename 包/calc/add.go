package calc

import (
	"fmt"
)
// go语言程序导入包语句会自动触发内部init()函数的调用
func init(){
	fmt.Println("calc_init")
}
// 标识符首字母大写表示对外可见
func Add(x, y int) int{
	return x + y
}