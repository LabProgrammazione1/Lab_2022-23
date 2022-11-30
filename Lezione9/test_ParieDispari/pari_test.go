package main

import (
	"testing"
)

func TestPariConDispari(t *testing.T) {
	if Pari(1) {
		t.Errorf("Pari(%d) = %v; atteso %v", 1, true, false)
	}
	if Pari(5) {
		t.Errorf("Pari(%d) = %v; atteso %v", 5, true, false)
	}
	if Pari(11) {
		t.Errorf("Pari(%d) = %v; atteso %v", 11, true, false)
	}
	if Pari(10000001) {
		t.Errorf("Pari(%d) = %v; atteso %v", 10000001, true, false)
	}
}

func TestPariConPari(t *testing.T) {
	if !Pari(6) {
		t.Errorf("Pari(%d) = %v; atteso %v", 6, false, true)
	}
	if !Pari(2) {
		t.Errorf("Pari(%d) = %v; atteso %v", 2, false, true)
	}
	if !Pari(10) {
		t.Errorf("Pari(%d) = %v; atteso %v", 10, false, true)
	}
	if !Pari(1000000) {
		t.Errorf("Pari(%d) = %v; atteso %v", 1000000, false, true)
	}
}

func TestPariConDispariNegativo(t *testing.T) {
	if Pari(-1) {
		t.Errorf("Pari(%d) = %v; atteso %v", -1, true, false)
	}
	if Pari(-5) {
		t.Errorf("Pari(%d) = %v; atteso %v", -5, true, false)
	}
	if Pari(-11) {
		t.Errorf("Pari(%d) = %v; atteso %v", -11, true, false)
	}
	if Pari(-10000001) {
		t.Errorf("Pari(%d) = %v; atteso %v", -10000001, true, false)
	}
}

func TestPariConPariNegativo(t *testing.T) {
	if !Pari(-6) {
		t.Errorf("Pari(%d) = %v; atteso %v", -6, false, true)
	}
	if !Pari(-2) {
		t.Errorf("Pari(%d) = %v; atteso %v", -2, false, true)
	}
	if !Pari(-10) {
		t.Errorf("Pari(%d) = %v; atteso %v", -10, false, true)
	}
	if !Pari(-1000000) {
		t.Errorf("Pari(%d) = %v; atteso %v", -1000000, false, true)
	}
}

func TestPariConZero(t *testing.T) {
	if !Pari(0) {
		t.Errorf("Pari(%d) = %v; atteso %v", 0, false, true)
	}
}
