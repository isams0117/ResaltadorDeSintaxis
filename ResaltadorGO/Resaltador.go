package main

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

// Códigos de color ANSI
const (
	ColorReset   = "\033[0m"
	ColorGreen   = "\033[92m"
	ColorYellow  = "\033[93m"
	ColorRed     = "\033[91m"
	ColorGray    = "\033[90m"
	ColorMagenta = "\033[95m"
	ColorCyan    = "\033[96m"
	ColorBlue    = "\033[94m"
)

type Token struct {
	File  string
	Type  string
	Value string
}

var tokenPatterns = map[string]string{
	"KEYWORD":    `\b(if|else|while|match|switch|case|return|lambda|of|end|fun)\b`,
	"LOGICAL":    `\b(and|or|not)\b|&&|\|\||!`,
	"COMPARISON": `==|!=|<=|>=|<|>|/=|=<`,
	"ARITHMETIC": `[+\-*/]`,
	"BOOLEAN":    `\b(true|false|True|False)\b`,
	"NULL":       `\b(None|nullptr)\b`,
	"NUMBER":     `\b\d+(\.\d+)?\b`,
	"STRING":     `"[^"]*"|'[^']*'`,
	"COMMENT":    `#.*|//.*|%.*`,
	"IDENTIFIER": `\b[a-zA-Z_][a-zA-Z0-9_]*\b`,
}

func colorize(tokenType string, value string) string {
	switch tokenType {
	case "KEYWORD":
		return ColorGreen + value + ColorReset
	case "LOGICAL":
		return ColorYellow + value + ColorReset
	case "ARITHMETIC", "COMPARISON":
		return ColorRed + value + ColorReset
	case "COMMENT":
		return ColorGray + value + ColorReset
	case "STRING":
		return ColorMagenta + value + ColorReset
	case "BOOLEAN":
		return ColorCyan + value + ColorReset
	case "NULL":
		return ColorBlue + value + ColorReset
	case "NUMBER":
		return ColorYellow + value + ColorReset
	default:
		return value
	}
}

// Cada archivo tiene su propio canal
func processFile(filename string, ch chan<- Token, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // cierra su canal al terminar

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("✗ %s no encontrado\n", filename)
		return
	}

	text := string(content)

	for tokenType, pattern := range tokenPatterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(text, -1)
		for _, match := range matches {
			ch <- Token{
				File:  filename,
				Type:  tokenType,
				Value: match,
			}
		}
	}
}

func main() {
	start := time.Now()

	files := []string{
		"Codigo1.py",
		"Codigo2.cpp",
		"Codigo3.erl",
	}

	// Crear un canal por archivo
	channels := make([]chan Token, len(files))
	var wg sync.WaitGroup

	for i, file := range files {
		channels[i] = make(chan Token, 100) // buffer de 100 tokens
		wg.Add(1)
		go processFile(file, channels[i], &wg) // goroutine por archivo
	}

	fmt.Println("TOKENS ENCONTRADOS")
	fmt.Println("==================")

	// Leer canales EN ORDEN — un archivo a la vez
	for i, ch := range channels {
		fmt.Printf("\n=== %s ===\n", files[i])
		for token := range ch {
			colored := colorize(token.Type, token.Value)
			fmt.Printf("  %-12s -> %s\n", token.Type, colored)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("\nTiempo de ejecución: %v\n", elapsed)
}