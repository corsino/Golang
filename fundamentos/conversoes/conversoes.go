package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado isso não converte int para string, apenas pega o valor equivalente do int
	fmt.Println("Teste " + string(97))

	//converte inteiro pra string
	fmt.Println("Teste " + strconv.Itoa(123))

	//convert string para inteiro

	num, _ := strconv.Atoi("123")

	println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		println("Verdadeiro")
	}

}
