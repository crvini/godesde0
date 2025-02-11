package maps

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	// fmt.Println(paises)
	// fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("equipo %s, tiene un puntajde de %d \n", equipo, puntaje)
	// }
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturea es %d, y el equipo existe = %t \n", puntaje, existe)
}
