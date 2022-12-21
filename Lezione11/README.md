# Laboratorio 11 - Metodi, Ricorsione

## 1) Metodi - Carte da gioco

Si consideri il seguente programma per la  rappresentazione e gestione di carte da gioco "francesi" (Esercizio1, Lezione 9).
Trasformare le funzioni `Mischia`, `Preleva` e `Riponi` del seguente codice in Metodi.

```go
package main

import (
	"fmt"
	"math/rand"
)

type CartaDaGioco struct {
	simbolo string
	seme    string
}

type Mazzo struct {
	carte  []CartaDaGioco
	ncarte int
}

var semi = [4]string{"quadri", "picche", "cuori", "fiori"}
var simboli = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}

// crea una carta da gioco: restituisce la carta da gioco dato il seme e il simbolo
func CreaCarta(simbolo, seme string) CartaDaGioco {
	return CartaDaGioco{simbolo, seme}
}

// crea e restituisce un mazzo di 40 carte da gioco
func CreaMazzo() Mazzo {
	tmp := make([]CartaDaGioco, 40)
	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			carta := CreaCarta(simboli[j], semi[i])
			tmp[k] = carta
			k++
		}
	}
	mazzo := Mazzo{carte: tmp, ncarte: 40}
	return mazzo
}

// mischia un mazzo di carte, usando il pacchetto "math/rand"
func Mischia(m Mazzo) {
	for i := range m.carte {
		j := rand.Intn(i + 1)
		m.carte[i], m.carte[j] = m.carte[j], m.carte[i]
	}
}

// simula il prelievo di una carta dalla cima di un mazzo;
func Preleva(m *Mazzo) (CartaDaGioco, error) {
	if m.ncarte > 0 {
		prelevata := m.carte[0]
		m.carte = m.carte[1:len(m.carte)]
		m.ncarte--
		return prelevata, nil
	} else {
		var empty_m Mazzo
		m = &empty_m
		return CartaDaGioco{"", ""}, fmt.Errorf("Mazzo Terminato")
	}
}

// simula la posa di una carta in cima ad un mazzo
func Riponi(m *Mazzo, c CartaDaGioco) error {
	if m.ncarte < 40 {
		m.carte = append([]CartaDaGioco{c}, m.carte...)
		m.ncarte++
		return nil
	} else {
		return fmt.Errorf("Mazzo Pieno")
	}
}

func main() {

	mazzo := CreaMazzo()

	fmt.Println(mazzo)
	fmt.Println()

	Mischia(mazzo)
	fmt.Println(mazzo)

	carta, _ := Preleva(&mazzo)
	fmt.Println()
	fmt.Println(carta)
	fmt.Println(mazzo)

	_ = Riponi(&mazzo, carta)
	fmt.Println()
	fmt.Println(mazzo)

	_ = Riponi(&mazzo, carta)
	fmt.Println()
	fmt.Println(mazzo)
}
```

## 2) Ricorsione - Potenze

La seguente funzione `Pow` calcola ricorsivamente `x` elevato a `n`

```text
x^0 = 1               --- caso base
x^n = x * x^(n-1)     
```

```go
package main

import "fmt"

func Pow(x float64, n uint) float64 {
	fmt.Println("n :", n)
	if n == 0 {
		return 1
	}
	return x * Pow(x, n-1) // calcola ricorsivamente x^(n - 1)
}

func main() {
	Pow(2,4)
}

```

## 2) Ricorsione - Potenze (2)

La stessa funzione può essere scritta in modo più efficiente:

```text
x^0 = 1               --- caso base
x^n = (x^(n/2))^2     se n > 0 e n è pari
x^n = x * x^(n-1)     se n > 0 e n è dispari
```

```go
package main

import "fmt"

func Pow(x float64, n uint) float64 {
	fmt.Println("n :", n)
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		halfpow := Pow(x, n/2) // calcola ricorsivamente x^(n/2)
		return halfpow * halfpow
	}
	return x * Pow(x, n-1) // calcola ricorsivamente x^(n - 1)
}

func main() {
	Pow(2,4)
}

```

## 3) Ricorsione - Fibonacci (1)

`Fib` calcola ricorsivamente un elemento della successione di Fibonacci.

- Soluzione ricorsiva semplice ma inefficiente: valori della successione sono calcolati (ricorsivamente) più volte.

Ad es. fib(5) è dato da fib(4) e fib(3), ma fib(4) ricalcola fib(3), il quale si basa su fib(2), che viene a sula volta ricalcolato più volte ...

```go
package main

import "fmt"

func Fib(x uint) uint64 {
	fmt.Println("Fib", x)
	if x < 2 {
		return uint64(x)
	}
	return Fib(x-1) + Fib(x-2)

}

func main() {
   Fib(4)
}

```

## 4) Ricorsione - Fibonacci (2)

`Fib2` è una versione ricorsiva alternativa, molto più efficiente.
Si basa su fib2 (vedi sotto), che sfrutta una mappa passata come argomento per memorizzare i valori della successione, non appena sono calcolati (memoization).
Si noti l'inizializzazione della mappa con i primi due valori della successione.

```go
package main

import "fmt"

func Fib2(x uint) uint64 {
	mfib := map[uint]uint64{0: 0, 1: 1} // inizialmente mfib contiene fib(0) e fib(1) (0,1 risp.)
	return fib2(x, mfib)
}

// fib2 calcola un elemento della successione, verificando preliminarmente
// che esso non sia già memorizzato nella mappa. Nel caso non lo sia, lo calcola (ricorsivamente)
// e lo memorizza.
func fib2(x uint, mfib map[uint]uint64) uint64 {
	fmt.Println("Fib (v.2)", x)
	if v, lookup := mfib[x]; lookup { // verifichiamo se fib(x) è stato già calcolato
		return v
	}
	mfib[x] = fib2(x-2, mfib) + fib2(x-1, mfib) // se fib(x) non è memorizzato, lo calcoliamo (ricorsivamente) e lo memorizziamo
	return mfib[x]
}

func main() {
   Fib2(10)
}

```

## 5) Ricorsione - Hanoi Towers

Implementazione ricorsiva del gioco delle "Hanoi Towers"; le tre torri (di partenza, di arrivo, di appoggio) sono denotate da tre etichette (stringhe): `from`, `to`, `spare`

```go
package main

import "fmt"

// stampa una mossa singola (spostamento di un singolo disco da from a to) -- definito per convenienza
func moveOneDisk(from, to string) {
	fmt.Println(from, "->", to)
}

// MoveDisks stampa ricorsivamente le singole mosse necessarie per spostare n dischi da from a to usando spare come supporto
func MoveDisks(n int, from, to, spare string) {
	
}

func main() {
  // illustra le mosse necessarie per spostare 4 dischi da "from" a "to" usando "spare"
  MoveDisks(4, "from", "to", "spare")
}

```

## 6) Ricorsione - Power Set

`PowSet` calcola l'insieme di tutti i sotto-insiemi di un insieme dato (usiamo insiemi di interi, ma potrebbero essere di qualsiasi altro tipo).

Ad es., se `A = {1,2,3,4}`, allora `powset(A)` è dato da:

```text
{}, {1}, {2}, {1,2}, {3}, {1,3}, {2,3}, {1,2,3}, {4}, {1,4}, {2,4}, {1,2,4}, {3,4}, {1,3,4}, {2,3,4}, {1,2,3,4} 
```

***NOTA*** Un insieme è rappresentato come slice di int ([]int) tuttavia l'ordine degli elementi nella slice, così come eventuali ripetizioni, non contano


```go
package main

import "fmt"

func PowSet(set []int) (pset [][]int) {
	if len(set) == 0 { // insieme vuoto (caso base)	
        return [][]int{ []int{} } // una slice contenente a una sola slice, la slice vuota --- caso base
       }
	 // passo(i) ricorsivo(i)
	 smaller := PowSet(set[:len(set)-1])  // calcoliamo RICORSIVAMENTE il power-set di set partendo da quello di set senza un elemento (l'ultimo)
	 for _, small := range smaller { // per ogni insieme (small) del power-set ottenuto da set senza il suo ultimo elemento
			large := []int{set[len(set)-1]} // creiamo un nuovo insieme (large) contenente l'ultimo elemento di set e small 
			large = append(large, small...)  //  (aggiungiamo l'intero contenuto della seconda slice alla prima)
			pset =  append(pset, large) // aggiungiamo l'insieme appena costruito (large) a pset
	  }
	  pset = append(pset, smaller...) // infine aggiungiamo a pset l'intero contenuto di smaller
        return pset
}

func main() {
   set := []int{1, 2, 3, 4}
   fmt.Println(PowSet(set))

}
```