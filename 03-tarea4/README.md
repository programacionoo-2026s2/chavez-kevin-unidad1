# Tarea 4 - Conversor organizado en paquetes# Tarea 4 - Conversor organizado en paquetes

## Descripción y estructura
Conversor de unidades de almacenamiento (`B`, `KB`, `MB`, `GB`) reorganizado en un paquete interno para separar la lógica de conversión del manejo de argumentos de consola.

```text
03-tarea4/
├── go.mod
├── main.go
├── README.md
└── internal/
    └── almacenamiento/
        └── almacenamiento.go

go.mod: define el módulo tarea4.

main.go: valida los argumentos de consola, llama al paquete almacenamiento y maneja los códigos de salida (ErrUso, ErrValor, ErrUnidad).

internal/almacenamiento/almacenamiento.go: contiene las constantes de bytes (B, KB, MB, GB), las funciones exportadas Bytes y Convertir, y la función privada redondear.

Cómo ejecutar
Compilar y ejecutar desde 03-tarea4:

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

Experimento del compilador
Al cambiar temporalmente la llamada en main.go a almacenamiento.redondear(valor) y ejecutar go build ./..., el compilador devolvió:
# tarea4
.\main.go:33:2: declared and not used: bytesOrigen
.\main.go:39:2: declared and not used: bytesDestino
.\main.go:45:30: name redondear not exported by package almacenamiento

Preguntas

a) ¿Por qué Bytes y Convertir empiezan con mayúscula y redondear con minúscula? ¿Qué ganas al dejar redondear oculta?
En Go, los nombres con mayúscula son exportados (públicos) y los que empiezan con minúscula son privados del paquete. Bytes y Convertir necesitan usarse desde main.go, mientras que redondear es solo un detalle interno de Convertir. Dejarla oculta evita que otros archivos dependan de ella y me permite cambiar cómo redondeo sin romper nada afuera.

b) ¿Qué mensaje dio el compilador al usar almacenamiento.redondear desde main.go y qué regla de Go lo explica?
Dio .\main.go:45:30: name redondear not exported by package almacenamiento. Esto pasa por la regla de visibilidad de Go: cualquier identificador que empiece con minúscula solo puede usarse dentro de su propio paquete.

c) ¿Para qué sirve que el paquete esté dentro de internal/ y qué pasaría si otro módulo intentara importarlo?
La carpeta internal/ protege el paquete para que solo pueda importarlo código dentro del mismo módulo (tarea4). Si otro módulo distinto intentara importar tarea4/internal/almacenamiento, el compilador de Go lo rechazaría con un error de compilación.

