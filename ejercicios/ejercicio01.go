package ejercicios

import (
	"strconv"
)

func ConvertidorEjercicio(numero string) (string, int) {
	numerConvert, err := strconv.Atoi(numero)
	if err != nil {
		return "Error: No se pudo convertir el número", 0
	}
	if numerConvert > 100 {
		return "Es Mayor a 100", numerConvert
	} else {
		return "Es menor a 100", numerConvert
	}
}
