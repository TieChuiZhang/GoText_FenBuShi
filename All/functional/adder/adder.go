//File  : adder.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-07-29

package mainXXX

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder2(0)
	//for i := 0; i<10; i++ {
	//	var s int
	//	s, a = a(i)
	//
	//	fmt.Printf("0+1+...+%d = %d\n",i,s)
	//}

	f := Fibonacci()
	printFileContents(f)
}

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
