package main

import (
	fm "fmt"
)
func main() {
	//print(111111111)
	//fm.Println("hello world")
	//sayHello22()
	//f := "short"
	//fm.Println(f)
	//for i:=1; i<5; i++ {
	//	fm.Print(i)
	//}
	//arr()

}

func arr() {
	//var a [3]int = [3]int{1, 2, 3}
	//fmt.Println(a)
	//for _, v := range a {
	//	fm.Printf("%d ", v)
	//}

	a := [...]    int{1, 2, 3}
	for _, v := range a {
		fm.Printf("%d", v)
	}

	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fm.Println(RMB, symbol[RMB])
}
