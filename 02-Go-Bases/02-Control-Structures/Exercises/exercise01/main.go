package main

import "fmt"

/*
=> Ejercicio 1
La Real Academia Española quiere saber cuántas letras tiene una palabra y
luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:

Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimir cada una de las letras.
*/

func main() {

	word := "application"

	// Count the letters of the word
	fmt.Println("The word", word, "has", len(word), "letters")

	// Print each letter of the word
	for i := 0; i < len(word); i++ {
		fmt.Println("Letter", i+1, "of the word", word, "is", string(word[i]))
	}

}
