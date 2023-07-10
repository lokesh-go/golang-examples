package main

import (
	_ "net/http/pprof"
	test "github.com/lokesh-go/golang-examples/pprof/test"
)

func main() {
	test.ComplexFunc()
	test.SimpleFunc()
}
