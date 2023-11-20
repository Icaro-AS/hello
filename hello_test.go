package main

import "testing"

func TestHello(t *testing.T){
	//Arrange
	esperado := "Olá, mundo!!!"
	
	//act
	resultado := Hello()
	
	//assert
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}