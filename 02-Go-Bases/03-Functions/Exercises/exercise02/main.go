package main

import "fmt"

/*

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
No se pueden introducir notas negativas.

*/

func getAverage(notes ...int) float64 {
	var sum float64
	for _, note := range notes {

		sum += float64(note)
	}
	average := sum / float64(len(notes))
	return average
}

func main() {
	average1 := getAverage(76, 68, 86, 92, 95, 81, 77, 51, 98, 98)
	fmt.Println("The average is", average1)

	average2 := getAverage(84, 56, 23, 98, 78, 45, 67, 89, 90, 100)
	fmt.Println("The average is", average2)
}
