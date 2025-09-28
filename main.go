package main

import (
	"fmt"
	"sync"
)

func main() {
	//task1
	//singleNumber([]int{2, 2, 1})
	//println(isValid("()[]{}"))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(plusOne([]int{1, 2, 3}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	//task2
	//指针
	//num := 10
	//recePointerAdd(&num)
	//fmt.Println(num)

	//a, b, c := 1, 2, 3
	//nums := []*int{&a, &b, &c}
	//recePointerMul(nums)
	//fmt.Println(a, b, c)

	//go print1_10()
	//go print2_10()
	//time.Sleep(time.Second * 1)

	//r := Rectangle{}
	//r.Area()
	//r.Perimeter()
	//c := Circle{}
	//c.Area()
	//c.Perimeter()

	//p := Person{"12", 12}
	//e := Employee{p, 54}
	//fmt.Println(e)

	//ch := make(chan int)
	//go sendNum(ch,10)
	//go receNum(ch,10)
	//time.Sleep(time.Second * 2)

	//ch2 := make(chan int, 100)
	//go sendNum(ch2, 100)
	//go receNum(ch2, 100)
	//time.Sleep(time.Second * 2)

	//var mu = sync.Mutex{}
	//count :=int32(0)
	//s := safeCounter{mu, count}
	//for i := 0; i < 10; i++ {
	//	go s.inc1000()
	//}
	//time.Sleep(time.Second * 5)
	//fmt.Println(s.count)

	var wg sync.WaitGroup
	a := atomicCounter{0}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go a.atomicInc1000(&wg)
	}
	wg.Wait()
	fmt.Println(a.count)
}
