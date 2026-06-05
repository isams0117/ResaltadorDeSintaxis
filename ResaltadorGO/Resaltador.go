package main 

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

//Almacena la información de cada token encontrado
type Token struct {
	File  string
	Type  string
	Value string
}

//Patrones a buscar dentro de los archivos
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


//Procesa el archivo completo
//results es el canal por donde se enviarán los tokens
func processFile(filename string, results chan<- Token, wg *sync.WaitGroup) {
	defer wg.Done()

	content, err := os.ReadFile(filename) //Lee todo el arhivo y se guarda como texto
	if err != nil {
		fmt.Printf("Error leyendo %s: %v\n", filename, err)
		return
	}

	text := string(content) //Convertir a string

	//Recorre todos los patrones
	for tokenType, pattern := range tokenPatterns {
		re := regexp.MustCompile(pattern) //crea el regex

		matches := re.FindAllString(text, -1) //busca coincidencias

		//Recorrer coincidencias
		for _, match := range matches {
			//envía el token por el canal
			results <- Token{
				File:  filename,
				Type:  tokenType,
				Value: match,
			}
		}
	}
}

func main() {

	start := time.Now() //marcar el inicio de la ejecución

	//archivos a analizar
	files := []string{
		"Codigo1.py",
		"Codigo2.cpp",
		"Codigo3.erl",
	}

	results := make(chan Token) //crear canal por donde viajarán los tokens

	var wg sync.WaitGroup //crear waitgroup 
	//Para contar cuantas goroutines siguen trabajando

	//crear go routines
	for _, file := range files {
		wg.Add(1) //incrementar contador
		go processFile(file, results, &wg) //lanzar goroutine
	}

	//cerrar canal cuando todos terminen
	go func() {
		wg.Wait() //esta goroutines espera hasta que las demás terminen
		close(results)
	}()

	fmt.Println("TOKENS ENCONTRADOS")
	fmt.Println("==================")

	//Consumir resultados
	//Escucha continuamente el canal y cada que llegue un token lo imprime
	for token := range results {
		fmt.Printf("[%s] %-8s -> %s\n",
			token.File,
			token.Type,
			token.Value)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nTiempo de ejecución: %v\n", elapsed)
}

