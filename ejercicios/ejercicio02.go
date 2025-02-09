package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func Multiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un número para multiplicar :")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	for i := 1; i <= 10; i++ {
		println("Multipilcación de ", numero1, " x ", i, " Es: ", numero1*i)
	}
}
