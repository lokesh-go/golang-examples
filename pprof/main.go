package main

import (
	_ "net/http/pprof"

	"github.com/lokesh-go/golang-examples/pprof/task"
)

func main() {
	task.ComplexFunc()
	task.SimpleFunc()
}
