package main

import "time"

func main() {
	println("String: ", "Banana" == "Banana")

	println("!=", 3 != 2)

	println("<", 3 < 2)

	println(">", 3 > 2)

	println("<=", 3 <= 2)

	println(">=", 3 >= 2)

	//datas
	d1 := time.Unix(0, 0)

	d2 := time.Unix(0, 0)
	println("Datas", d1 == d2)

	println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	println("Pessoas", p1 == p2)

}
