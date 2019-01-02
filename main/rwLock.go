package main

import (
	"sync"
	"fmt"
	"math/rand"
)

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func main() {

	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}
}

func read (n int)  {
	rw.RLock()
	fmt.Printf("%d reading ...",n)

	v := count

	fmt.Printf("%d readed count value %d",n,v)
	wg.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("%d writing ...",n	)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("%d writed count value %d",n,count)
	wg.Done()
	rw.RUnlock()
}