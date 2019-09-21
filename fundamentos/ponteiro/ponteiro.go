package main

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o endereço de memória da variavel i
	*p++   //desreferenciando (pegando o valor)
	i++

	//Go não tem aritmetica de ponteiros
	//p++

	println(p, *p, i, &i)
}
