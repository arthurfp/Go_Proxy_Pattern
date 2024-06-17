package main

import (
	"fmt"
	"proxy-pattern-go/pkg/proxy"
)

func main() {
	realSubject := &proxy.RealSubject{}
	proxy := proxy.NewProxy(realSubject)

	fmt.Println(proxy.Request())
}
