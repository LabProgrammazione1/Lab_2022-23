package main

import (
	"testing"
)

func TestDispariConDispari(t *testing.T) {
	if !Dispari(1) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 1, false, true)
	}
	if !Dispari(5) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 5, false, true)
	}
	if !Dispari(11) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 11, false, true)
	}
	if !Dispari(10000001) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 10000001, false, true)
	}
}

func TestDispariConPari(t *testing.T) {
	if Dispari(6) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 6, true, false)
	}
	if Dispari(2) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 2, true, false)
	}
	if Dispari(10) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 10, true, false)
	}
	if Dispari(1000000) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 1000000, true, false)
	}
}

func TestDispariConDispariNegativo(t *testing.T) {
	if !Dispari(-1) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -1, false, true)
	}
	if !Dispari(-5) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -5, false, true)
	}
	if !Dispari(-11) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -11, false, true)
	}
	if !Dispari(-10000001) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -10000001, false, true)
	}
}

func TestDispariConPariNegativo(t *testing.T) {
	if Dispari(-6) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -6, true, false)
	}
	if Dispari(-2) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -2, true, false)
	}
	if Dispari(-10) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -10, true, false)
	}
	if Dispari(-1000000) {
		t.Errorf("Dispari(%d) = %v; atteso %v", -1000000, true, false)
	}
}

func TestDispariConZero(t *testing.T) {
	if Dispari(0) {
		t.Errorf("Dispari(%d) = %v; atteso %v", 0, true, false)
	}
}
