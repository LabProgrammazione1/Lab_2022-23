package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		f.Close()
		return
	}

	for {
		var nome string
		var x, y float64
		_, err = fmt.Fscan(f, &nome, &x, &y)
		if err == io.EOF {
			f.Close()
			break
		}
		if err != nil {
			fmt.Printf("Error while reading the file! %v\n", err)
			f.Close()
			return
		}
		fmt.Printf("Punto %s = (%v, %v)\n", nome, x, y)
	}
}