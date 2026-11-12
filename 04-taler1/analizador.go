package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

	totalBytes := len(texto)
	totalRunas := utf8.RuneCountInString(texto)

	palabras := strings.Fields(texto)
	totalPalabras := len(palabras)

	mayusculas := 0
	digitos := 0
	conTildeONie := 0
	letrasEspeciales := "áéíóúÁÉÍÓÚñÑ"

	for _, r := range texto {
		if unicode.IsUpper(r) {
			mayusculas++
		}
		if unicode.IsDigit(r) {
			digitos++
		}
		if strings.ContainsRune(letrasEspeciales, r) {
			conTildeONie++
		}
	}

	fmt.Printf("Bytes: %d\n", totalBytes)
	fmt.Printf("Runas: %d\n", totalRunas)
	fmt.Printf("Palabras: %d\n", totalPalabras)
	fmt.Printf("Mayúsculas: %d\n", mayusculas)
	fmt.Printf("Dígitos: %d\n", digitos)
	fmt.Printf("Con tilde/ñ: %d\n", conTildeONie)

}
