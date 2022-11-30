package main

import "fmt"

func main() {
	fmt.Print("Inserisci un numero: ")
	var n int
	fmt.Scan(&n)

	if Pari(n) {
		fmt.Printf("Il numero %d è pari\n", n)
	} else {
		fmt.Printf("Il numero %d è dispari\n", n)
	}

}

func Pari(p int) bool {
	return p%2 == 0
}

func Dispari(n int) bool {
	return !Pari(n)
}
