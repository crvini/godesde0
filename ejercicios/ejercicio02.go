package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func Multiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número para multiplicar :")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, i*numero1)
	}
	return texto
}
