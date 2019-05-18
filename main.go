package main

import "fmt"
import "./hyperdimentional"

func main() {
	fmt.Println("hello world")

	x := hyperdimentional.NewHdVec()

    x.Print()

	fmt.Scanln()
}
