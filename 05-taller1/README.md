# Taller 1 - Estadísticas de notas

## Descripción
Programa en Go con una función variádica `estadisticas` que recibe cualquier cantidad de notas (`...float64`) y devuelve la nota menor, la mayor, el promedio y un error centinela (`ErrSinDatos`) si no se pasan valores.

## Cómo ejecutar
Desde la carpeta `05-taller1`:

```powershell
go run .


Qué imprime cada línea:

Línea 1: llamada con valores sueltos (8.5, 6, 9.75, 4), mostrando menor, mayor, promedio y <nil>.

Línea 2: llamada pasando un slice con notas..., mostrando min, max y prom con 2 decimales.

Línea 3: llamada sin argumentos que devuelve ErrSinDatos y se valida con errors.Is.


Pruebas realizadas

4 9.75 7.0625 <nil>
min=7.00 max=10.00 prom=8.17
error: no se recibieron datos

Preguntas

a) ¿Por qué los resultados se llaman menor y mayor y no min y max? Explica qué pasaría con la llamada min(menor, n) si el resultado se llamara min.
Porque Go ya tiene funciones integradas llamadas min y max. Si el resultado se llamara min, esa variable float64 taparía la función integrada y al intentar llamar min(min, n) Go daría un error de compilación (cannot call non-function min (variable of type float64)), porque intentaría usar una variable numérica como si fuera una función.

b) ¿Qué hacen los tres puntos en estadisticas(notas...)? ¿Qué mensaje te da Go si los quitas y por qué?
Los tres puntos (...) desempaquetan el slice notas para pasar cada nota como un argumento float64 suelto a la función variádica. Si los quitas (estadisticas(notas)), Go da un error de compilación (cannot use notas (variable of type []float64) as float64 value in argument to estadisticas) porque la función espera valores float64 individuales y no un slice completo.