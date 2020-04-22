package main

import (
	"fmt"
	"time"
)

//并发测试demo
func main() {
	//demo1 goroutine协程测试
	go routine1("这是一个在协程中执行的函数")
	routine1("这是一个正常执行的函数")

	//demo2 channel通道测试
	td := make(chan int)
	//声明一个切片数组
	b := []int{1, 2, 3, 4, 5, 8, 9, 10, 11}
	//切片其中的后半部分进行计算和
	go sum(b[:len(b)/2], td)
	//切片前半部分进行计算和
	go sum(b[len(b)/2:], td)

	d, e := <-td, <-td
	fmt.Println(d, e, d+e)

}

func sum(s []int, td chan int) {
	fmt.Println(s)
	//计算后存储的和
	sum := 0
	//遍历s数组，并存储至v变量中
	for _, v := range s {
		//sum+v后的值存储至sum中
		sum += v
	}
	//把计算后的结果存储到通道中
	td <- sum
}

//goroutine协程测试
func routine1(b string) {

	for i := 0; i < 5; i++ {
		//随机延迟等待时间
		time.Sleep(10 * time.Millisecond)
		//打印传入的字符串
		fmt.Println(b, i)
	}
}
