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

	menor, mayor = nums[0], nums[0]
	suma := 0.0

	for _, n := range nums {
		suma += n
		menor = min(menor, n)
		mayor = max(mayor, n)
	}

	promedio = suma / float64(len(nums))
	return
}

func main() {
	fmt.Println(estadisticas(8.5, 6, 9.75, 4))

	notas := []float64{10, 7, 7.5}
	mn, mx, pr, _ := estadisticas(notas...)
	fmt.Printf("min=%.2f max=%.2f prom=%.2f\n", mn, mx, pr)

	if _, _, _, err := estadisticas(); errors.Is(err, ErrSinDatos) {
		fmt.Println("error:", err)
	}
}
