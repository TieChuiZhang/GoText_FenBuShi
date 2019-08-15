//File  : main.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-07-29

package main

import (
	mainXXX "Object/FenBuShi/All/functional/adder"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("start")
	try()
	writeFile("fib.txt")
}

func try() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)

}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
		fmt.Println("error:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	f := mainXXX.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
