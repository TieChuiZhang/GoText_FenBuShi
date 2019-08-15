//File  : add_test.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-15

package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("start")
}
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}
	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle (%d, %d);"+
				"got %d; expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
