//File  : queue_test.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-20

package queue

import (
	"fmt"
)

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	// Output:
	// 2
	// 1
	// 2
	// false
	// 3
	// true
	// want:
	// 2
	// FAIL

}
