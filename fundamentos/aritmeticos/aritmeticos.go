package main

import "math"

func main() {
	a := 3
	b := 2
	println("Soma => ", a+b)

	println("subtração => ", a-b)

	println("Divisão => ", a/b)

	println("Multiplicação => ", a*b)

	println("Módulo => ", a%b)

	//bitwise
	println("E=>", a&b) // 11 & 10  = 10

	println("Ou=>", a|b) // 11 | 10  = 11

	println("Xor=>", a^b) // 11 ^ 10  = 01
	c := 3.0
	d := 2.0

	//outras operações usando math
	println("Maior =>", math.Max(float64(a), float64(b)))
	println("Menor =>", math.Min(c, d))
	println("Exponenciação =>", math.Pow(c, d))
}
