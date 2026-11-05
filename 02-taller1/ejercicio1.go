package main

import "fmt"

var a = b + 1
var b = valor("b", 10)

func valor(nombre string, v int) int {

	fmt.Println("Inicializando", nombre)
	return v
}

func init() { fmt.Println("init 1: a = ", a) }
func init() { fmt.Println("init 2") }
func main() {
	fmt.Println("main: a =", a, "b =", b)
}
