package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Inserisci testo (termina con CTRL+D):")

	scanner := bufio.NewScanner(os.Stdin)

	testo := ""
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}

	fmt.Print("Testo letto:\n", testo)

}