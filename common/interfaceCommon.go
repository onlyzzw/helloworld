package common

import "fmt"

func NewLoginer() Loginer {
	return defaultLogin(0)
}

type defaultLogin int

type Loginer interface {
	Login()
}


func (d defaultLogin) Login(){
fmt.Println("login in...")
}


