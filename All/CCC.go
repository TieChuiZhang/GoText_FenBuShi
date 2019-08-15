package main

import (
	"Object/FenBuShi/All/en/leemock"
	leereal "Object/FenBuShi/All/en/real"
	"fmt"
	"time"
)

func main1() {
	//en.Abc()
	var r Retriever
	retriever := leemock.Retriever{"this is fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &leereal.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)
	if realRetriever, ok := r.(*leereal.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *leemock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *leereal.Retriever:
		fmt.Println("Contents", v.UseraAgent)
	}
}

type Retriever interface {
	Get(url string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "goland",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}
