package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error while opening the file! %v\n", err)
		f.Close()
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}