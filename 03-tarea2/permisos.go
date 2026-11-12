package main

import "fmt"

type Permiso uint8

const (
	Ejecutar Permiso = 1 << iota
	Escribir
	Leer
)

func describir(p Permiso) string {
	r, w, x := "-", "-", "-"
	if p&Leer != 0 {
		r = "r"
	}
	if p&Escribir != 0 {
		w = "w"
	}
	if p&Ejecutar != 0 {
		x = "x"
	}
	return r + w + x
}

func main() {
	usuario := Leer | Escribir | Ejecutar
	grupo := Leer | Ejecutar
	otros := Leer

	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))
	fmt.Printf("numérico: %d%d%d\n", usuario, grupo, otros)

	otros |= Escribir
	grupo &^= Ejecutar

	fmt.Println(describir(usuario) + describir(grupo) + describir(otros))
}
