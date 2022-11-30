package main

import (
	"testing"
	"reflect"
)

func TestSt(t *testing.T) {
	if !reflect.DeepEqual(Sottostringhe("123456"), []string{"12", "123", "1234", "12345", "123456", "23", "234", "2345", "23456", "34", "345", "3456", "45", "456", "56"}){
		t.Errorf("Errore 1!")
	}
	if len(Sottostringhe("654321")) != 0 {
		t.Errorf("Errore 2!")
	}
	if !reflect.DeepEqual(Sottostringhe("123121212"), []string{"12", "123", "23", "12", "12", "12"}) {
		t.Errorf("Errore 3!")
	}
	if !reflect.DeepEqual(Sottostringhe("01010101"), []string{"01", "01", "01", "01"}) {
		t.Errorf("Errore 4!")
	}
}

func TestTesto(t *testing.T) {
	if len(Sottostringhe("acc23")) != 0 {
		t.Errorf("Errore Testo!")
	}
}

