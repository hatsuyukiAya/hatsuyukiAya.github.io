package main

import "fmt"

func main() {
	//数组声明方式
	//方式1
	type num [3]int
	var arr1 num
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	//方式2
	var nums = [3]int{1,2,3}
	nums[1] = 4


	fmt.Println(nums)

	//方式3
	// * 这样可以不用自己确认长度，但这种只能是在静态初始化的时候
	var nums2 = [...]int{1,2,3}
	//var nums2 = [...]int  这样也会报错 由于没有内容 无法确认长度
	fmt.Println(nums2,"\t", len(nums2))
	fmt.Println(arr1,"arr1")
	//获取数组长度
	fmt.Println(len(nums))

}
