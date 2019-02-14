package main

import (
	"fmt"
	"./mock"
	"./real"
)

type Retriever interface {
	Get(url string) string

}

func download(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}

func main()  {
	var r Retriever
	r = mock.Retriever{"这是接口实现"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))

}
