package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crvini/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Multiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo ", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Multiplicar()

	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Eerror durante el append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Eerror durante el writestring" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchvio() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error leyendo archivo", err.Error())
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
