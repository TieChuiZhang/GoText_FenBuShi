//File  : main.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-27

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	encoding2 "golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}
	// gbk 转utf-8
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}

func determineEncoding(
	r io.Reader) encoding2.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
