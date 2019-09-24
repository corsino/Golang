package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i <= 10 { //não tem while em go
		println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Printf("%d", i)
	}

	println("\nMisturando...")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			print("Par ")
		} else {
			print("Impar ")
		}
	}

	for {
		// laço infinito
		println("Para sempre...")
		time.Sleep(time.Second)
	}
}
