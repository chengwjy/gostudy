package main

import "fmt"
import "strings"
func main() {
	// 数组
	// 数组是同一种数据类型元素的集合
	// 数组从声明时就确定，使用时可以修改数组成员，但是
	// 数组的大小不可变化
	 // 数组的长度必须是常量，并且长度是数组类型的一部分
	var a [3]int 
	var b [2]int
	//a = b // 不可以这样做，因为此时a和b是不同类型
	fmt.Println(a, b)
	// 数组的初始化
	// 1.定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海"} // 最后两个元素为空字符串
	cityArray[len(cityArray) - 1] = "深圳"
	fmt.Println(cityArray)
	// 2.编译器推导数组的长度
	var boolArray = [...]bool{true, false, true}
	fmt.Println(boolArray)
	// 3.使用索引值的方式初始化
	var langArray = [...]string{1: "golang", 3: "C++"}
	fmt.Println(len(langArray)) // 4
	// 遍历数组
	for index, value := range langArray{
		fmt.Println(index, value)
	}
	// 二维数组
	// 二位数组只有外层size可以用...表示
	var school = [...][2]string{
		{"初一3班", "初一4班"},
		{"初三1班", "初三5班"},
	}
	for _, grade := range school{
		for _, classname := range grade{
			fmt.Println(classname)
		}
		fmt.Println("grade——end")
	}
	// GO语言中数组是值类型，不是引用类型

	// 切片
	// 切片是一个引用相同类型元素的可变长度的序列
	// 切片是一个引用类型,是对数组的封装
	// 它包含了三个信息:底层数组的指针、切片的长度、容量
	var qp = []string{"十月", "清晨"}
	fmt.Println(qp)

	// 基于数组得到切片
	var names = [...]string{"tom", "lucy", "jack"}
	// 0：从第0个键开始（包括
	// 2：到第二个键结束（不包括
	var someName = names[0:2]
	// 切片再切片 [:] 相当于 [0:len]
	fmt.Println(someName[:]) // [tom lucy]
	// make构造切片 5：元素个数 10：容量
	productQp := make([]int, 5, 10)
	fmt.Println(productQp)
	// 通过cap()函数获取切片的容量
	fmt.Println(cap(productQp))
	// 切片之间不能比较，我们不能用==操作符来判断两个切片是否含有全部想等元素
	// 切片唯一合法的比较操作是和nil比较
	// 注意：len(切片) == 0 和 切片 == nil 两者是有区别的
	fmt.Println(productQp == nil) // false
	// 使用append() 函数为切片添加元素
	// 当底层数组不能容纳新增的元素时，切片会自动按照一定的策略
	// 进行扩容，此时该切片只想的底层数组就会更换。
	productQp = append(productQp, 998, 997)
	fmt.Println(productQp)
	// append()函数追加切片
	nums := []int{3, 2, 1}
	productQp = append(productQp, nums...)
	fmt.Println(productQp)
	// 切片的copy。由于直接赋值切片会发生多个变量引用同一个切片地址的问题，copy应运而生
	qpA := []int{1,2,3}
	qpB := make([]int, 2, 2)
	copy(qpB, qpA)
	// qpB容量只有2，所以只copy了qpA的前2个元素
	fmt.Println(qpB) // [1 2]
	// 删除切片的元素 删除键为1的元素李白
	authors := []string{"毛泽东", "李白", "杜甫"}
	authors = append(authors[:1], authors[2:]...)
	fmt.Println(authors)

	// map
	// go语言中提供的映射关系容器map,其内部使用散列表(hash)实现
	// make初始化 ;注意：这里,3是容量
	spaces := make(map[string]string, 3)
	// 添加键值对
	spaces["鲁迅"] = "周树人"
	fmt.Println(spaces)
	// 直接赋值
	var testMap = map[string]string{"江西": "jiangxi", "浙江": "zhejiang"}
	fmt.Println(testMap)
	// 判断台湾是否在testMap里面
	// _ 接收值，如果不存在则是该基本类型的默认值。ok接收布尔值，是否存在
	_, ok := testMap["台湾"]
	fmt.Println(ok)
	// map是无序的，键值对和添加的顺序无关
	for k, v := range testMap{
		fmt.Println(k, v)
	}

	// 统计一个字符串每个单词出现的次数
	word := "how do you do"
	wordQp := strings.Split(word, " ")
	wordMap := make(map[string]int)
	for _, v := range(wordQp){
		_, ok := wordMap[v]
		if(ok){
			wordMap[v]++
		}else{
			wordMap[v] = 1
		}
	}	
	fmt.Println(wordMap)

} 