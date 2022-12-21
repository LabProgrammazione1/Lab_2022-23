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