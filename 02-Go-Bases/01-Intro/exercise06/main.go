package main

func main() {
	var num1 int = 10
	var num2 float64 = 3.5

	// Type convertion
	var result float64 = float64(num1) + num2
	var result2 = float64(num1) * num2
	var result3 = num1 * int(num2)

	println(result, result2, result3)

}
