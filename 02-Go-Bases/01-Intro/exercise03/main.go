package main

import "fmt"

func main() {
	// Declara variables para el título de un libro (string), año de publicación (int), y si lo has leído o no (bool).
	title := "Clean code"
	year := 2010
	read := true

	fmt.Printf("The book %s was published in %d and I have read it: %t\n", title, year, read)
}
