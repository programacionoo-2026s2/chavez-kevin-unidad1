04 - Taller 1: Analizador de Texto en Go

¿Qué hace el programa?

Bytes y runas: Cuenta los bytes con len() y los caracteres (runas) con utf8.RuneCountInString().

Palabras: Separa y cuenta las palabras usando strings.Fields().

Mayúsculas y dígitos: Recorre el texto con range y evalúa cada carácter con unicode.IsUpper() y unicode.IsDigit().

Tildes y ñ: Detecta vocales con tilde o la letra ñ usando strings.ContainsRune().

¿Por qué los bytes (62) y las runas (58) no coinciden?

En Go los textos usan codificación UTF-8:

1. Las letras normales, números y espacios ocupan 1 byte.

2. Los 4 caracteres especiales del texto (ó, á, ¡, ñ) ocupan 2 bytes cada uno.

3. Por eso hay 58 caracteres (runas), pero en memoria suman 62 bytes (58 + 4 bytes extra = 62).

Ejecución y resultado

go run .


Bytes: 62
Runas: 58
Palabras: 9
Mayúsculas: 3
Dígitos: 4
Con tilde/ñ: 3