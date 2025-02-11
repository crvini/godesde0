package funciones

import "fmt"

func Exponencial(valor int) {
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)
	Exponencial(valor * 2)
}
