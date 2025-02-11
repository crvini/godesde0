package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)

	posicion := arreglo[3:]   // Slice creado desde un vector, desde la posición 3
	posicion2 := arreglo[:5]  // Slice creado desde un vector , desde la posición 0 hasta la 5
	posicion3 := arreglo[6:7] // Slice creado desde un vector, desdee la posición 6 hasta la 7

	fmt.Println(posicion)
	fmt.Println(posicion2)
	fmt.Println(posicion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, capacidad %d", len(nums), cap(nums))
}
