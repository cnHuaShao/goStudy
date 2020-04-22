package main

import "fmt"

//主方法
func main() {
	demo2()
}

//多种方式打印
func demo1() {

}

//数据特性测试
func demo2() {
	//声明一个int类型的数组，并赋值
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("声明的a数组的指针是：%x \n", a)
	demo2_son1(a)
	fmt.Printf("传入数组时的值：%x \n", a)
	demo2_son2(&a)
	fmt.Printf("传入数组指针时的值：%x \n", a)

	//声明一个切片，初始存储一个长度为3的int类型数组，默认值为[0,0,0]，声明最大扩展长度为5，
	slice1 := make([]int, 3, 5)
	demo2_slice1(slice1)
	fmt.Printf("传入数组时的值：%x \n", slice1)
	demo2_slice2(slice1)
}

//传入整个数组，底层值完全复制出来了一份，新的内存地址
func demo2_son1(b [5]int) {
	//打印当前数组的内存地址
	fmt.Printf("当前传入数组的指针是：%x \n", b)
	//改变数组中第二个索引下的值
	b[2] = 8
	fmt.Println("函数体内值：", b[2])

}

//传入整个数组的值，不一样的指针地址，指针指向同一个值存储的内存地址。
func demo2_son2(b *[5]int) {

	//打印当前数组的内存地址
	fmt.Printf("当前传入数组的指针是：%x \n", b)
	//改变数组中第二个索引下的值
	b[2] = 8
	fmt.Println("函数体内值：", b[2])
}

//传入整个数组，
func demo2_slice1(b []int) {
	//打印当前数组的内存地址
	fmt.Printf("当前传入数组的指针是：%x \n", b)
	//改变数组中第二个索引下的值
	b[2] = 8
	fmt.Println("函数体内值：", b[2])

}
func demo2_slice2(b []int) {
	fmt.Println("当前切片底层的数组长度是：", len(b)) // 3
	fmt.Println("当前切片最大可扩展长度是：", cap(b)) // 5
}
