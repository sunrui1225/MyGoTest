package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
	err := returnsError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}

//1）returnsError() 函数不返回 error 非空接口类型，而是直接返回结构体指针 *MyError（明确的类型，阻止自动装箱）；
//2）不要直接 err != nil 这样判断，而是使用类型断言来判断：
//if e, ok := err.(*MyError); ok && e != nil {
//fmt.Printf("error occur: %+v\n", e)
//return
//}
