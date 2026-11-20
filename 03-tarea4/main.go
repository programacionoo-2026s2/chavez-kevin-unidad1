package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"tarea4/internal/almacenamiento"
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

	origen := strings.ToUpper(os.Args[2])
	destino := strings.ToUpper(os.Args[3])

	bytesOrigen, existe := almacenamiento.Bytes(origen)
	if !existe {
		fmt.Fprintf(os.Stderr, "Error: unidad de origen no reconocida: %s\n", origen)
		os.Exit(ErrUnidad)
	}

	bytesDestino, existe := almacenamiento.Bytes(destino)
	if !existe {
		fmt.Fprintf(os.Stderr, "Error: unidad de destino no reconocida: %s\n", destino)
		os.Exit(ErrUnidad)
	}

	resultado := almacenamiento.Convertir(valor, bytesOrigen, bytesDestino)
	fmt.Printf("%g %s = %.2f %s\n", valor, origen, resultado, destino)
}
