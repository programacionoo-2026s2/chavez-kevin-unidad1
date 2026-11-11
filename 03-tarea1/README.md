1. ¿Qué imprime? 
total: 0
2. ¿Por qué ocurre? (Explicación del «ámbito»)
El resultado total: 0 ocurre debido al ensombrecimiento de variables (variable shadowing) provocado por el operador de declaración corta := dentro del bucle for:

    1. Ámbito de la función (suma): Al inicio de la función, la instrucción total := 0 declara e inicializa la variable total perteneciente al ámbito local de la función.

    2. Ámbito del bloque (for): Dentro del cuerpo del bucle for, la línea total := total + i utiliza nuevamente :=. En lugar de reasignar el valor a la variable existente, Go declara una nueva variable local con el mismo nombre (total), cuyo ciclo de vida está restringido exclusivamente al ámbito del bloque interno del bucle.

    3. Efecto de ensombrecimiento: La variable del ámbito interno oculta (ensombrece) a la variable total del ámbito superior. Los incrementos se realizan sobre la variable temporal interna (y la línea _ = total evita que el compilador reclame por variable no utilizada). Cada vez que una iteración finaliza, esa variable interna sale de su ámbito y se destruye.

    4. Retorno: Al llegar a return total, se evalúa la variable correspondiente al ámbito de la función, la cual nunca fue modificada y conserva su valor inicial 0.

3. Corrección cambiando un solo carácter
Para corregir el error y que el programa sume correctamente (1 + 2 + 3 = 6), se debe cambiar un solo carácter: eliminar los dos puntos (:) en la asignación dentro del bucle:

Línea ordinaria: total := total + i

Línea corregida: total = total + i

Explicación (Al cambiar el operador de declaración corta := por el operador de asignación simple =, ya no se crea una nueva variable en el ámbito del bloque. En su lugar, se actualiza directamente el valor de la variable total declarada en el ámbito de la función y ahora sale el total de "6")

4. Extra: Ejecución de go vet ./... ¿Lo detecta?
No, go vet ./... no lo detecta por defecto. Pero claramente tiene un explicacion del porque no lo detecta por defecto y se lo explicare acontinuacion: La herramienta estándar go vet ejecuta un conjunto predefinido de analizadores considerados de alta precisión y bajo nivel de falsos positivos (por ejemplo: formato de Printf, structs mal inicializados, llamadas atómicas inválidas, etc.). Sin embargo, el análisis de ensombrecimiento (shadowing) está desactivado por defecto en la suite básica de go vet para evitar exceso de advertencias en código idiomático donde el ensombrecimiento suele ser intencional (como al reutilizar err := ...).
5. Investigación del analizador shadow
shadow es un analizador de código estático oficial mantenido por el equipo de Go, alojado en el paquete golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow. Su función exclusiva es inspeccionar el código fuente e identificar declaraciones de variables que oculten o ensombrezcan variables homónimas declaradas en un bloque o ámbito exterior.
