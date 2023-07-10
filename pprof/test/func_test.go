package test

import "fmt"

func SimpleFunc() {
	println("Simple")
}

func ComplexFunc() {
	for i := 0; i < 10e5; i++ {
		go print(i)
	}
}

func print(i int) {
	fmt.Printf("%d  ", i)
}
