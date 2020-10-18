package main

import "fmt"
import "encoding/json"
// import "strings"

// 匿名结构体
type address struct{
	city string
}
type person struct{
	// 结构体的匿名字段
	string
	int
	// 结构体的匿名嵌套
	address
}

// 结构体的继承
type Animal struct{
	name string
}
func (a *Animal)move(){
	fmt.Printf("%s 会动\n", a.name)
}
type Dog struct{
	Feet int8
	*Animal
}


type class struct{
	Title string
	Students []string
}

func main() {
	p1 := person{
		"小王子",
		18,
		address{
			"上海",
		},
	}
	fmt.Println(p1)
	// 通过类型访问
	fmt.Println(p1.string)
	fmt.Println(p1.address.city)
	// 可以直接访问它嵌套的结构体的属性
	fmt.Println(p1.city)

	d1 := Dog{
		4,
		&Animal{
			"贝贝",
		},
	}
	fmt.Println(d1)
	// 结构体嵌套实现继承
	d1.move()

	// 结构体字段的可见性与JSON序列化
	// 结构体中字段大写开头表示可公开访问
	// 小写表示私有（仅在定义当前结构体的包中可访问）

	c1 := class{
		Title: "三年二班",
		Students: []string{"李红", "李白"},
	}
	fmt.Println(c1)

	// json序列化：Go语言中的数据-> JSON格式的字符串
	// 得到的是[]byte类型
	data, err := json.Marshal(c1)
	fmt.Println(string(data))
	fmt.Println(err)

	// 反序列化
	jsonStr := "{\"Title\":\"三年二班\",\"Students\":[\"李红\",\"李白\"]}"
	var c2 class
	err2 := json.Unmarshal( []byte(jsonStr), &c2)

	fmt.Println(c2)
	fmt.Println(err2)
}
