package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	fmt.Printf("newSlice长度:%d,容量:%d",len(newSlice),cap(newSlice))
	fmt.Println()

	newSlice[0] = 10
	fmt.Println(newSlice)
	fmt.Println(slice)


	newSlice=append(newSlice,10)
	fmt.Println(newSlice)
	fmt.Println(slice)
	fmt.Printf("newSlice长度:%d,容量:%d",len(newSlice),cap(newSlice))
	fmt.Println()
	fmt.Printf("slice长度:%d,容量:%d",len(slice),cap(slice))
	fmt.Println()

	fmt.Printf("%p\n", &slice)
	modify(slice)
	fmt.Println(slice)
	fmt.Println(newSlice)
}

func modify(slice []int)  {
	fmt.Printf("%p\n",slice)
	slice[1]=100

}
