# Tarea 3 - Conversor de unidades

## Descripción
Programa de consola en Go que convierte valores entre unidades de almacenamiento (`B`, `KB`, `MB`, `GB`) usando múltiplos de 1024, recibe los datos por línea de comandos y devuelve distintos códigos de salida si ocurre un error.

## Cómo ejecutar
Compilar y ejecutar desde la carpeta `03-tarea3`:

```powershell
go build -o convertir.exe .
.\convertir.exe <valor> <unidad_origen> <unidad_destino>


Códigos de salida:

0: conversión exitosa.

1 (ErrUso): cantidad incorrecta de argumentos.

2 (ErrValor): el valor no es un número válido o es negativo.

3 (ErrUnidad): unidad de origen o destino no reconocida.

Pruebas realizadas

Comando

convertir 2048 KB MB
convertir 1 gb mb
convertir 1500 B KB
convertir 0.5 MB KB
convertir abc KB MB
convertir -5 KB MB
convertir 5 TB MB
convertir 5 KB

Salida obtenida

2048 KB = 2.00 MB
1 GB = 1024.00 MB
1500 B = 1.46 KB
0.5 MB = 512.00 KB
Error: el valor debe ser un número mayor o igual a 0
Error: el valor debe ser un número mayor o igual a 0
Error: unidad de origen no reconocida: TB
Uso: convertir <valor> <unidad_origen> <unidad_destino>

Código

0
0
0
0
2
2
3
1

Preguntas

a) ¿Por qué el valor está en os.Args[1] y no en os.Args[0]?
Porque os.Args[0] siempre guarda el nombre o la ruta del ejecutable (convertir.exe). Los datos que escribe el usuario empiezan en os.Args[1], por eso ahí va el valor.

b) ¿Por qué los mensajes de error se escriben en stderr y no con fmt.Println?
Porque stderr separa los errores de la salida normal (stdout). Así, si otro programa o script lee el resultado de mi conversor, los mensajes de error no se mezclan con el dato convertido.

c) ¿Qué pasaría en tu programa si escribieras valor, _ := strconv.ParseFloat(...) e ignoraras el error al ejecutar convertir abc KB MB?
Go asignaría 0 a valor sin detenerse, el programa imprimiría 0 KB = 0.00 MB y terminaría con código 0, ocultando que "abc" no era un número.

d) ¿Qué ventaja tiene declarar los códigos de salida con const en lugar de escribir 1, 2 y 3 dentro de os.Exit?
El código es más claro al leer os.Exit(ErrUnidad) en vez de un número suelto, y si después necesito cambiar un código solo lo edito una vez en el bloque const.