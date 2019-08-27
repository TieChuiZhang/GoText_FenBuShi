//File  : regex.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-27

package main

import (
	"fmt"
	"regexp"
)

const text = `My email is cc@xxxx.com
LiLei email is   cc@xxxx.com
ZhangHua email    is cc@xxxx.com
HanHan email    is cc@xxxx.com.cn
`

func main() {
	// .+ .* 点匹配任何字符加号就是一个或者多个星号就是一个或者多个
	// \\. \\匹配转义字符 当前匹配. 如不需要转义需要使用单斜引号``
	// [a-zA-Z0-9]
	// 提取正则中的数据加() FindAllStringSubmatch
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
