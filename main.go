package main

import (
	"fmt"

	"github.com/crvini/godesde0/variables"
)

func main() {
	entero, texto := variables.ConviertoaTexto(15885)
	fmt.Println(entero)
	fmt.Print(texto)
}
