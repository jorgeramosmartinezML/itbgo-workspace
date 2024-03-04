package main

import "fmt"

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:

Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:

Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que
retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:

const (
   dog    = "dog"
   cat    = "cat"
)

...

animalDog, msg := animal(dog)
animalCat, msg := animal(cat)

...

var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func dogFood(amount int) float64 {
	return float64(amount) * 10
}

func catFood(amount int) float64 {
	return float64(amount) * 5
}

func hamsterFood(amount int) float64 {
	return float64(amount) * 0.25
}

func tarantulaFood(amount int) float64 {
	return float64(amount) * 0.15
}

func animal(animal string) (func(int) float64, string) {
	switch animal {
	case dog:
		return dogFood, ""
	case cat:
		return catFood, ""
	case hamster:
		return hamsterFood, ""
	case tarantula:
		return tarantulaFood, ""
	default:
		return nil, "Animal not found"
	}
}

func main() {

	animalDog, _ := animal(dog)
	animalCat, _ := animal(cat)

	amount := 10
	var kg float64
	kg = animalDog(10)
	fmt.Println("The amount of dog food is", kg, "kg for", amount, "dogs")

	amount = 12
	kg = animalCat(amount)
	fmt.Println("The amount of cat food is", kg, "kg for", amount, "cats")

}
