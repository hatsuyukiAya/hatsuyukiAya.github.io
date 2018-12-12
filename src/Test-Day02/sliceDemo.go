package main

import "fmt"

func main() {
	/*
	   切片(Slice)
	   跟java的ArrayList相似，可变数组
	   切片是一种引用数据类型，底层是一个数组的引用，所以在改变引用值的时候，切片的值也会改变

	 */
	//声明方式
	type ms []int;
	var a ms
	fmt.Println(a)


	var ms1 = []int{1,2,3}

	fmt.Println(ms1)
	//添加一个或者多个元素，返回相同类型的Slice
	ms1 = append(ms1,5,6,1)
	//获得Slice长度 len
	println("长度为",len(ms1))
	//获得容器最大长度
	println("最大长度为",cap(ms1))

	//把数组切片给另一个数组
	var number3 = []int{1,2,3,4,5}
	var slice1  = number3[1:4]	// [2 3 4] 包前不包后
	fmt.Println(slice1)

}
