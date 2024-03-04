package exercise04

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Ejercicio 4 - Testear el cálculo de estadísticas
Los profesores de la universidad de Colombia, entraron al programa
de análisis de datos  de Google, el cual premia a los mejores
estadísticos de la región. Por ello los profesores nos solicitaron
comprobar el correcto funcionamiento de nuestros cálculos estadísticos.
Se solicita la siguiente tarea:

Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.
*/

func TestOperation(t *testing.T) {
	t.Run("Calculate min value", func(t *testing.T) {
		// Arrange
		operation := "minimum"
		numbers := []int{1, 2, 3, 4, 5}
		expectedResult := 1

		// Act
		funMin, err := Operation(operation)
		result := funMin(numbers...)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("Calculate max value", func(t *testing.T) {
		// Arrange
		operation := "maximum"
		numbers := []int{1, 2, 3, 4, 5}
		expectedResult := 5

		// Act
		funMax, err := Operation(operation)
		result := funMax(numbers...)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("Calculate average value", func(t *testing.T) {
		// Arrange
		operation := "average"
		numbers := []int{1, 2, 3, 4, 5}
		expectedResult := 3

		// Act
		funAverage, err := Operation(operation)
		result := funAverage(numbers...)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("Operation not found", func(t *testing.T) {
		// Arrange
		operation := "operation not found"

		// Act
		_, err := Operation(operation)

		// Assert
		require.Error(t, err)
	})
}
