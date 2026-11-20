package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	ErrUso    = 1
	ErrValor  = 2
	ErrUnidad = 3
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || valor < 0 {
		fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número mayor o igual a 0")
		os.Exit(ErrValor)
	}

	fmt.Println("Valor válido:", valor)
}
