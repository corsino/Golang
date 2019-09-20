package main

import "fmt"

func main() {

	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" Nova")
	fmt.Println(" Linha")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("O valo de x é " + xs)

	fmt.Println("O valo de x é", x)

	// Imprime formatado
	fmt.Printf("O valo de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
