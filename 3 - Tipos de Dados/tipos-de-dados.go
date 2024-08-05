package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	// FIM NUMEROS INTEIROS

	//NÚMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 127473487493.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)
	//FIM NUMEROS REAIS

	//STRINGS
	var str string = "ABCDE"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)
	//FIM STRINGS

	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
