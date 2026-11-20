package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	LimiteSubtitulo = 42
	ErrUso          = 1
	ErrVacio        = 2
	ErrLimite       = 3
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `Uso: analizar "<texto>"`)
		os.Exit(ErrUso)
	}

	texto := os.Args[1]
	if strings.TrimSpace(texto) == "" {
		fmt.Fprintln(os.Stderr, "Error: el texto no puede estar vacío ni tener solo espacios")
		os.Exit(ErrVacio)
	}

	fmt.Printf("texto: %q\n", texto)
}
