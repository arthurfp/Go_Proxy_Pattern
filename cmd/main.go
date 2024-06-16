package main

import (
	"fmt"
	"proxy-pattern-go/pkg/proxy"
)

func main() {
	realSubject := &proxy.RealSubject{}
	fmt.Println(realSubject.Request())
}
