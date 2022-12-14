package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("tabellina.txt")
	if err != nil {
		fmt.Printf("Error while creating the file! %v\n", err)
		f.Close()
		return
	}

	for i:=1; i<=10; i++ {
		risultato := i * 10
		_, err = fmt.Fprintln(f, "10 x ", i, risultato)
		if err != nil {
			fmt.Printf("Error while writing the file! %v\n", err)
			f.Close()
			return
		}
	}
	f.Close()
}