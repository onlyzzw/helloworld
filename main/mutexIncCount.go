package main

import (
	fm "fmt"
	"sync"
	"runtime"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
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

	//互斥
	mutexIncCount()

}

func mutexIncCount() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fm.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}

