package main

import "fmt"

/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
calificaciones de los/as estudiantes de un curso.
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.

Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere
realizar (mínimo, máximo o promedio) y que devuelva otra función y un mensaje
(en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y
devuelva el cálculo que se indicó en la función anterior.

Ejemplo:

const (
   minimum = "minimum"
   average = "average"
   maximum = "maximum"
)

...

minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...

minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Min(numbers ...int) int {
	minValue := numbers[0]
	for _, number := range numbers {
		if number < minValue {
			minValue = number
		}
	}
	return minValue
}

func Max(numbers ...int) int {
	maxValue := numbers[0]
	for _, number := range numbers {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue
}

func Average(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

func operation(operation string) (func(...int) int, error) {
	switch operation {
	case minimum:
		return Min, nil
	case average:
		return Average, nil
	case maximum:
		return Max, nil
	default:
		return nil, fmt.Errorf("operation %s not found", operation)
	}
}

func main() {

	minFunc, err1 := operation(minimum)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("The minimum value is", minValue)

	}

	averageFunc, err2 := operation(average)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("The average value is", averageValue)
	}

	maxFunc, err3 := operation(maximum)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("The maximum value is", maxValue)

	}

	unknownFunc, err4 := operation("unknown")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		unknownValue := unknownFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Println("The unknown value is", unknownValue)
	}

}
