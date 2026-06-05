# ResaltadorDeSintaxis

Analizador concurrente de código Go que identifica y extrae tokens (palabras clave, números y comentarios) de múltiples archivos en paralelo.

## 📋 Descripción

**ResaltadorDeSintaxis** es un programa escrito en Go que procesa archivos de código y detecta patrones específicos utilizando expresiones regulares. Utiliza goroutines y canales para procesar múltiples archivos de forma concurrente, aprovechando el paralelismo de Go.

## ✨ Características

- **Análisis Concurrente**: Procesa múltiples archivos de forma paralela usando goroutines
- **Detección de Patrones**:
  - **KEYWORD**: Identifica palabras clave (`if`, `else`, `while`, `return`, `switch`, `case`)
  - **NUMBER**: Detecta números enteros y decimales
  - **COMMENT**: Reconoce comentarios (formato `#` y `//`)
- **Comunicación Asíncrona**: Utiliza canales para sincronización entre goroutines
- **Thread-Safe**: Implementa `sync.WaitGroup` para coordinación segura

## 🏗️ Estructura del Código

### Tipos Principales

```go
type Token struct {
    File  string  // Nombre del archivo procesado
    Type  string  // Tipo de token (KEYWORD, NUMBER, COMMENT)
    Value string  // Contenido del token encontrado
}
```

### Patrones de Búsqueda

Los patrones se definen como expresiones regulares en el mapa `tokenPatterns`:

| Tipo | Patrón | Ejemplos |
|------|--------|----------|
| KEYWORD | `\b(if\|else\|while\|return\|switch\|case)\b` | `if`, `else`, `while` |
| NUMBER | `\b\d+(\.\d+)?\b` | `42`, `3.14` |
| COMMENT | `(#.*\|//.*)` | `# comentario`, `// comentario` |

### Funciones

- **`processFile(filename string, results chan<- Token, wg *sync.WaitGroup)`**: Procesa un archivo individual en una goroutine separada, extrae tokens y los envía por el canal
- **`main()`**: Orquesta la ejecución, crea goroutines para cada archivo y consume los resultados

## 🚀 Cómo Usar

### Requisitos
- Go 1.13 o superior

### Instalación y Ejecución

```bash
# Clonar el repositorio
git clone https://github.com/isams0117/ResaltadorDeSintaxis.git
cd ResaltadorDeSintaxis

# Ejecutar el programa
go run Resaltador.go
```

### Configuración

Para analizar diferentes archivos, modifica la sección `files` en la función `main()`:

```go
files := []string{
    "tu_archivo1.py",
    "tu_archivo2.cpp",
    "tu_archivo3.txt",
}
```

## 📤 Salida

El programa imprime los tokens encontrados en el siguiente formato:

```
TOKENS ENCONTRADOS
==================
[codigo1.py] KEYWORD  -> if
[codigo2.cpp] NUMBER   -> 42
[codigo3.erl] COMMENT  -> # este es un comentario
```

## 🔧 Personalización

### Agregar Nuevos Patrones

Edita el mapa `tokenPatterns` para incluir nuevos tipos de tokens:

```go
var tokenPatterns = map[string]string{
    "KEYWORD":   `\b(if|else|while|return|switch|case)\b`,
    "NUMBER":    `\b\d+(\.\d+)?\b`,
    "COMMENT":   `(#.*|//.*)`,
    "STRING":    `"[^"]*"`,  // Nuevo patrón
    "VARIABLE":  `\$\w+`,     // Nuevo patrón
}
```

## 🎯 Conceptos Go Utilizados

- **Goroutines**: Ejecución concurrente de funciones
- **Canales**: Comunicación segura entre goroutines
- **sync.WaitGroup**: Sincronización y espera de goroutines
- **Expresiones Regulares**: Búsqueda de patrones con `regexp`

## 📝 Licencia

Este proyecto está disponible bajo la licencia que el autor especifique.

## 🤝 Autor

isams0117