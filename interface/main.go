package main

import (
	"fmt"
	"./mock"
	"./real"
	"./People"
)

type Retriever interface {
	Get(url string) string

}

func download(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}

type Peoples interface {
	Eat(thing string) string
}

func xiaoming(p Peoples) string {
	return p.Eat("馒头")
}

func main()  {
	var r Retriever
	r = mock.Retriever{"这是接口实现"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))

	rr := People.China{}
	fmt.Println(xiaoming(rr))

}
