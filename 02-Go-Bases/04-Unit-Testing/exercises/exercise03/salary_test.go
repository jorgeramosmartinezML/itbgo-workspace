package exercise03

/*
Ejercicio 3 - Test del salario
La empresa marinera no está de acuerdo con los resultados obtenidos en
los cálculos de los salarios, por ello nos piden realizar una serie de
tests sobre nuestro programa. Necesitaremos realizar las siguientes
pruebas en nuestro código:

Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSalary(t *testing.T) {
	t.Run("Calculate salary of category A", func(t *testing.T) {
		// Arrange
		minutesWorked := 60
		category := "A"
		expectedSalary := 4500.0

		// Act
		salary := CalculateSalary(minutesWorked, category)

		// Assert
		require.Equal(t, expectedSalary, salary)
	})

	t.Run("Calculate salary of category B", func(t *testing.T) {
		// Arrange
		minutesWorked := 60
		category := "B"
		expectedSalary := 1800.0

		// Act
		salary := CalculateSalary(minutesWorked, category)

		// Assert
		require.Equal(t, expectedSalary, salary)
	})

	t.Run("Calculate salary of category C", func(t *testing.T) {
		// Arrange
		minutesWorked := 60
		category := "C"
		expectedSalary := 1000.0

		// Act
		salary := CalculateSalary(minutesWorked, category)

		// Assert
		require.Equal(t, expectedSalary, salary)
	})

	t.Run("Calculate salary of unknown category", func(t *testing.T) {
		// Arrange
		minutesWorked := 60
		category := "D"
		expectedSalary := 0.0

		// Act
		salary := CalculateSalary(minutesWorked, category)

		// Assert
		require.Equal(t, expectedSalary, salary)
	})
}
