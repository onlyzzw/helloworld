package main

import "fmt"

func main() {
	ad:=admin{user{"张三","zhangsan@flysnow.org"},"管理员"}
	sayHello(ad.user)//使用user作为参数
	sayHello(ad)//使用admin作为参数
}

type user struct {
	name string
	email string
}

type admin struct {
	user
	level string
}

type Hello interface {
	hello()
}

func (u user) hello(){
	fmt.Println("Hello，i am a user")
}

func sayHello(h Hello){
	h.hello()
}
