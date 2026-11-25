package main

import (
	"errors"
	"fmt"
)

var ErrSinDatos = errors.New("no se recibieron datos")

func estadisticas(nums ...float64) (menor, mayor, promedio float64, err error) {
	if len(nums) == 0 {
		err = ErrSinDatos
		return
	}
	return
}

func main() {
	fmt.Println(estadisticas())
}
