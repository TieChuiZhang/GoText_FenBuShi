//File  : recover.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-14

package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurrde", err)
		} else {
			//panic(r)
			panic(fmt.Sprintf("I don't know what to do:%v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	//b := 0
	//a := 5/b
	//fmt.Println(a)
	panic(123)

}
func main() {
	tryRecover()
}
