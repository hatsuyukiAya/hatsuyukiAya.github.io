package main

import (
	"fmt"
)

func main() {
	//Go的数据类型
	/*
	基本类型
	bool 布尔
	整型 int8 byte int16 int uint uintptr等
	浮点 float32 float64
	复数类型 complex64 complex128
	字符串 string
	字符类型 rune
	错误类型 error

	复合类型
	指针 pointer
	数组 array
	切片 slice
	字典 map
	通道 chan
	结构体 struct
	接口 interface

	*/

	//布尔
	var v1 bool
	v1 = true
	v2 := (1==2)
	fmt.Println(v1,v2)
	//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	//以下正确用法
	var b bool
	b = (1!=0)
	fmt.Println("Result:",b)	//输出结果 Result:true

	//整型
	/*
	int8  1  						-128 ~ 127(java的byte)
	uint8（即byte）  1 				 0 ~ 255
	int16  2  					-32768 ~ 32767(java的short)
	uint16  2  						0 ~ 65535
	int32  4  					-2147483648 ~ 2147483647(java的int)
	uint32  4  						0 ~ 4294967295
	int64  8  			-9223372036854775808 ~ 9223372036854775807(相当于java的long)
	uint64  8  				0 ~ 18446744073709551615
	int  平台相关  平台相关
	uint  平台相关  平台相关
	uintptr  同指针  在32位平台下为4字节，64位平台下为8字节
	*/
	//需要注意的是， int 和 int32 在Go语言里被认为是两种不同的类型，编译器也不会帮你自动
	//做类型转换，比如以下的例子会有编译错误
	/*
	var value2 int32
	value1 := 64 // value1将会被自动推导为int类型
	value2 = value1 // 编译错误
	*/
	//使用强制类型转换能解决这个编译错误
	var value2 int32
	value1 := 64
	value2 = int32(value1)	//强转
	fmt.Println(value2)
}
