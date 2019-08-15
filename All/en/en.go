package en

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
)

func Abc() {
	fmt.Println(1)
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}
