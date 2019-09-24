package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //executa na sequencia o próximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1:
		return "E"
	default:
		return "Nota inválida"

	}
}

func notaParaConceitoDesafio(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "A" //executa na sequencia o próximo case
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))

	fmt.Println(notaParaConceito(6.7))

	fmt.Println(notaParaConceito(2.1))

	fmt.Println(notaParaConceito(-9.8))
}
