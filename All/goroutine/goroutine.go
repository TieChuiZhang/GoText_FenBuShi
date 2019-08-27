//File  : goroutine.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-21

package main

import (
	"fmt"
	"time"
)

func main() {
	//var a[10] int
	for i := 0; i < 10; i++ {
		go func(i int) { //
			for {
				//a[i]++
				//runtime.Gosched()
				fmt.Println("hello from"+"goroutine %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}

func printHello() {

}
