package main
import (
	"fmt"
	// 当你的代码都放在$GOPATH目录下的话
	// 我们导入包的路径要从$GOPATH/src/后面继续写起
	// 以mycalc作为包的别名
	// 如果只希望导入包，而不是用包内部的数据，可以使用匿名导入包
	// import _ "包的路径"
	mycalc "包/calc"
)
func main(){
	result := mycalc.Add(1, 2)
	fmt.Println(result)
}