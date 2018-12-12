package main

import  "fmt"

func main()  {
	//变量的定义方式
	var(
		name string = "ichika"
		age int32 = 5025
		lockOn bool = true
	)
	var v1 int = 12
	var v2 = 15
	v3:= 10

	//变量数据交换
	v3,v2=v2,v3

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(v1,v2,v3)
	fmt.Println(lockOn)


	fmt.Println(GetName())

	//常量
	//const关键字可以给常量指定一个友好的名字
	const Pi float64 = 3.1415926535
	const zero = 0.0	//无类型浮点常量
	const(
		size int64 = 1024
		eof = -1	//无类型整型常量
	)
	const u , v float32 = 0,3	//u=0.0 v=3.0 常量的多重赋值
	const a,b,c = 3,4,"foo"		//a=3 b=4 c="foo" 无类型整型和字符串常量

	//总结
	//变量是运行中可以改变的值 需要用var修饰
	//常量不可在运行中修改 用const修饰
	/*
	由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达
	式，比如试图以如下方式定义常量就会导致编译错误：
	const Home = os.GetEnv("HOME")
	os.GetEnv("HOME")在运行中才能得出结果，无法在编译中确认值，所以无法作为常量定义的右值
	*/

	//预定义常量
	//Go语言预定义了这些常量： true 、 false 和 iota
	//iota比较特殊，再每一个const出现的时候会被重置为0，每出现一次iota，值加1
	const(
		c1 = iota
		c2 = iota
		c3 = iota
	)
	const(	//出现了const 重置const计数
		d1 = iota
		d2 = iota
		d3 = iota
	)
	fmt.Println(c1,c2,c3)
	fmt.Println(d1,d2,d3)

	//枚举
	//枚举指一系列相关的常量，比如下面关于一个星期中每天的定义
	//Go语言并不支持众多其他语言明确支持的 enum 关键字
	const(
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thurday
		Friday
		Saturday
		numberOfDays	//这个常量没有导出
	)
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thurday,Friday,Saturday)


}

func GetName() (firstName,lastName,nickName string){
	return "shiroha","haili","shiro"
}
