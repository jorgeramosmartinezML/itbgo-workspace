package exercise05

import "testing"

/*
Ejercicio 5 - Calcular cantidad de alimento

El refugio de animales envió una queja ya que el cálculo total de alimento
a comprar no fue el correcto y compraron menos alimento del que necesitaban.
Para mantener satisfecho a nuestro cliente deberemos realizar los test
necesarios para verificar que todo funcione correctamente:

Verificar el cálculo de la cantidad de alimento para los perros.
Verificar el cálculo de la cantidad de alimento para los gatos.
Verificar el cálculo de la cantidad de alimento para los hamster.
Verificar el cálculo de la cantidad de alimento para las tarántulas.
*/
func TestDog(t *testing.T) {
	// Arrange
	amount := 10
	pet := "dog"
	expectedResult := 100.0

	// Act
	fnDog, _ := animal(pet)
	result := fnDog(amount)

	// Assert
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestCat(t *testing.T) {
	// Arrange
	amount := 10
	pet := "cat"
	expectedResult := 50.0

	// Act
	fnCat, _ := animal(pet)
	result := fnCat(amount)

	// Assert
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestHamster(t *testing.T) {
	// Arrange
	amount := 10
	pet := "hamster"
	expectedResult := 2.5

	// Act
	fnHamster, _ := animal(pet)
	result := fnHamster(amount)

	// Assert
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestTarantula(t *testing.T) {
	// Arrange
	amount := 10
	pet := "tarantula"
	expectedResult := 1.5

	// Act
	fnTarantula, _ := animal(pet)
	result := fnTarantula(amount)

	// Assert
	if result != expectedResult {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func TestAnimal_NotFound(t *testing.T) {
	// Arrange
	pet := "notfound"
	expectedResult := "Animal not found"

	// Act
	_, result := animal(pet)

	// Assert
	if result != expectedResult {
		t.Errorf("Expected %s, got %s", expectedResult, result)
	}
}
