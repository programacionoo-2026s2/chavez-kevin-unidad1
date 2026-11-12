# 03-tarea2: Permisos con operadores de bits en Go

Programa `permisos.go` que representa permisos de estilo Unix (lectura, escritura y ejecución) usando operadores de bits.

1. Cómo ejecutarlo

Dentro de la carpeta `03-tarea2`:

```
go run .
```

Salida:

```
rwxr-xr--
numérico: 754
rwxr--rw-
```

2. Estrategia de desarrollo

1.1 Representar los permisos como bits

Se define `type Permiso uint8` y tres constantes con `iota`. Cada permiso ocupa un bit distinto, de modo que se pueden combinar sin que se pisen entre sí:

| Permiso   | Valor | Binario |
|-----------|-------|---------|
| Leer      | 4     | 100     |
| Escribir  | 2     | 010     |
| Ejecutar  | 1     | 001     |

`iota` arranca en 0 y aumenta en 1 por cada línea del bloque `const`. Con la fórmula `1 << iota` se obtiene 1, 2 y 4 en el orden en que están escritas las constantes.

1.2 Función `describir` (operador `&`)

Para saber si un permiso incluye una capacidad se usa el AND de bits (`&`). Por ejemplo, `p & Leer != 0` es verdadero solo si el bit de lectura está encendido en `p`. Según el resultado se arma el texto: la letra (`r`, `w` o `x`) si está encendido, o `-` si no.

3. Usuario, grupo y otros

- Usuario: `Leer | Escribir | Ejecutar` → `rwx`
- Grupo: `Leer | Ejecutar` → `r-x`
- Otros: `Leer` → `r--`

El operador `|` (OR de bits) combina permisos. Se imprimen los tres textos juntos: `rwxr-xr--`.
 
4. Valor numérico (`%d`)

Con `%d` se imprime el valor de cada permiso. Con el orden final de las constantes, usuario = 4+2+1 = 7, grupo = 4+1 = 5 y otros = 4, por lo que se imprime `754`.

5. Modificar permisos (`|=` y `&^=`)

- `otros |= Escribir` enciende el bit de escritura sin tocar los demás. Otros pasa de `r--` a `rw-`.
- `grupo &^= Ejecutar` apaga el bit de ejecución (AND NOT, "bit clear"). El grupo pasa de `r-x` a `r--`.

Resultado final: `rwxr--rw-`.

6. Reto: orden de las constantes como en `chmod`

El código de partida define `Leer = 1`, `Escribir = 2` y `Ejecutar = 4`. Con ese orden, el número impreso es `751` (7 + 5 + 1), que coincide con el ejemplo del enunciado.

Para que los valores sean los mismos que usa `chmod` (`r=4`, `w=2`, `x=1`) solo se cambió el orden en el que se declaran las constantes: `Ejecutar` primero, `Escribir` en el medio y `Leer` al final. Como `iota` asigna los valores según la posición, no hizo falta tocar ninguna otra línea del programa, porque el resto del código usa los nombres y no los números.

Con este orden el valor numérico es `754`, que es el mismo número que se obtendría con `chmod 754`. Las letras impresas no cambian; solo cambia el número.