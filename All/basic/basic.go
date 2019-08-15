//File  : basic.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-15

package main

import (
	"fmt"
	"math"
)

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
