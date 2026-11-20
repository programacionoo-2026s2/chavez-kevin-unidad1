# Taller 1 - FizzBuzz

## Descripción
Programa en Go que recorre los números del 1 al 30 usando un bucle `for` y un `switch` sin condición para imprimir `FizzBuzz` (múltiplo de 15), `Fizz` (múltiplo de 3), `Buzz` (múltiplo de 5) o el número, y al final muestra cuántas veces apareció cada palabra.

## Cómo ejecutar
Desde la carpeta `04-taller1`:

```powershell
go run .

Resultado
Salida del programa original:

1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 FizzBuzz
Fizz=8 Buzz=4 FizzBuzz=2

Salida del experimento (poniendo case i%3 == 0 antes de case i%15 == 0):

1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 Fizz 16 17 Fizz 19 Buzz Fizz 22 23 Fizz Buzz 26 Fizz 28 29 Fizz
Fizz=10 Buzz=4 FizzBuzz=0

Preguntas

a) ¿Qué salida y qué conteos obtuviste al poner el case de 3 antes que el de 15, y por qué FizzBuzz dejó de aparecer?
En el 15 y el 30 se imprimió Fizz en lugar de FizzBuzz, y los conteos cambiaron a Fizz=10 Buzz=4 FizzBuzz=0. FizzBuzz dejó de aparecer porque todo múltiplo de 15 también es múltiplo de 3, así que el programa siempre entraba primero al case i%3 == 0 y nunca llegaba al case i%15 == 0.

b) Cuando i vale 15, ¿cuántos case son verdaderos y cuál ejecuta Go? ¿Por qué no hace falta escribir break al final de cada case?
Son verdaderos 3 casos (i%15 == 0, i%3 == 0 y i%5 == 0), pero Go evalúa de arriba hacia abajo y ejecuta solo el primero que cumple (case i%15 == 0). No hace falta escribir break porque en Go cada case termina automáticamente al acabar su bloque, sin seguir hacia el siguiente.
