# Laboratorio 7 - Stringhe, Strutture, Puntatori

## (Stringhe) Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	s := "ciao, come va?" 
    /* s è interamente definita da caratteri considerati nello standard US-ASCII */

	fmt.Println(string(s[0]) + string(s[len(s)-2]) + string(s[len(s)-1]))
}

```

## (Funzioni con output multipli, errori) Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Errore, divisione per zero!")
	}
	return a/b, nil
}
```

## (Stringhe) Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	s := "ciao, come va?"
    /* s è interamente definita da caratteri considerati nello standard US-ASCII */

    fmt.Println(s[6:10] + string(s[len(s)-1]))
	fmt.Println(s[:5] + s[len(s)-4:])
	
}
```

## (Strutture) Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	var p1 Persona
	p1 = Persona{"Rick", "Sanchez"}

	var p2 Persona
	p2.Nome = "Jerry"
	p2.Cognome = "Smith"

	p3 := Persona{Nome: "Morty"}

	p2 = p3

	fmt.Println(p1, p2)
}
```

## (Puntatori) Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
	var a, b, c int
	var ptr *int
	a, b, c = 10, 15, 20
	ptr = &a
	*ptr += 10
	a += 10
	ptr = &b
	*ptr += 10
	*ptr = c
	*ptr += 10
	fmt.Println(a, b, c)
}
```

## (Puntatori) Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a, b = 15, 20
	f(&a, &b)
	fmt.Println(a, b)
}
func f(x, y *int) {
	*x, *y = *y, *x
}
```

## (Puntatori) Trova l'errore

Questo programma dovrebbe stampare `30 20` ma non genera l'output desiderato. Corregere e verificare che l'esecuzione del programma generi l'output atteso.

```go
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr *int
	ptr = a
	*ptr = *ptr + b
	fmt.Println(a, b)
}
```

## (Puntatori) Trova l'errore

Questo programma dovrebbe stampare `70 20` ma non genera l'output desiderato. Corregere e verificare che l'esecuzione del programma generi l'output atteso.

```go
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr *int
	*ptr = 50
	ptr = &a
	*ptr += b
	fmt.Println(a, b)
}
```

## (Puntatori) Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import (
	"fmt"
	"strings"
)

type Persona struct {
	Nome, Cognome string
}

func main() {

	p := &Persona{"Rick", "Sanchez"}

	f(*p)
	fmt.Println(*p)
	g(p)
	fmt.Println(*p)
}

func f(p Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}

func g(p *Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}
```