# Tarea 5 - Analizador de subtítulos


## Descripción
Programa de consola en Go que analiza un texto contando sus bytes, runas, palabras, letras, dígitos, mayúsculas y runas de más de 1 byte, y verifica si cumple con el límite accesible de 42 runas para subtítulos.

## Cómo ejecutar
Compilar y ejecutar desde `03-tarea5`. El texto debe ir entre comillas para que la terminal lo pase como un solo argumento:

```powershell
go build -o analizar.exe .
.\analizar.exe "Canción de cuna 2026"

Códigos de salida:

0: texto válido y dentro del límite de 42 runas.

1 (ErrUso): cantidad incorrecta de argumentos (por ejemplo, sin texto o sin comillas).

2 (ErrVacio): el texto está vacío o solo tiene espacios.

3 (ErrLimite): el texto supera las 42 runas.


Preguntas
a) En Canción de cuna 2026 tu programa muestra 21 bytes y 20 runas. ¿Por qué no coinciden y cuál de los dos números usarías para el límite de subtítulos?
No coinciden porque en UTF-8 la letra ó ocupa 2 bytes aunque sea un solo carácter, mientras que las demás ocupan 1 byte. Para el límite de subtítulos usaría las runas (20), porque quien lee el video ve caracteres en pantalla, no bytes.

b) En for i, r := range texto, ¿qué valen i y r en cada vuelta? Si texto es ñu, ¿qué valores toma i?
En cada vuelta, i es la posición en bytes donde empieza ese carácter y r es la runa (el carácter). Si texto es "ñu", i toma los valores 0 y 2, porque ñ ocupa 2 bytes y la u empieza en el byte 2.

c) ¿Por qué las palabras se cuentan con la variable enPalabra y no contando los espacios? Usa el caso 4 como ejemplo.
Porque si el texto tiene varios espacios seguidos o espacios al principio y al final, contar espacios da un número falso. En el caso 4 ("  espacios   múltiples  "), hay 7 espacios en total, pero usando enPalabra el programa cuenta correctamente 2 palabras.
