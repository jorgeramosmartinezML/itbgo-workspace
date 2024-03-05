package main

/*
Ejercicio 2 - Imprimir datos
A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer”
que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("It was not possible to read the file")
		}

		fmt.Println("Ejecución finalizada")
	}()

	data := readFile("customers.txt")
	fmt.Printf(data)

}

func readFile(nameFile string) (data string) {
	file, err := os.Open(nameFile)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	data = string(bytes)
	return

}
