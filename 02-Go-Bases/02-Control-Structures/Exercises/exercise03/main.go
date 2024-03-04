package main

import "fmt"

/*
=> Ejercicio 3
Realizar una aplicación que reciba  una variable con el número del mes.

Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.

Nota: Validar que el número del mes, sea correcto.
*/

func main() {

	// Option 1
	// I would choose this option because it's the most readable and it's the most efficient.
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	month := 7
	if month >= 1 && month <= 12 {
		fmt.Println(month, ",", months[month-1])
	} else {
		fmt.Println("Invalid month")
	}

	// Option 2
	// month = 7
	// switch month {
	// case 1:
	// 	fmt.Println(month, ", January")
	// case 2:
	// 	fmt.Println(month, ", February")
	// case 3:
	// 	fmt.Println(month, ", March")
	// case 4:
	// 	fmt.Println(month, ", April")
	// case 5:
	// 	fmt.Println(month, ", May")
	// case 6:
	// 	fmt.Println(month, ", June")
	// case 7:
	// 	fmt.Println(month, ", July")
	// case 8:
	// 	fmt.Println(month, ", August")
	// case 9:
	// 	fmt.Println(month, ", September")
	// case 10:
	// 	fmt.Println(month, ", October")
	// case 11:
	// 	fmt.Println(month, ", November")
	// case 12:
	// 	fmt.Println(month, ", December")
	// default:
	// 	fmt.Println("Invalid month")

}
