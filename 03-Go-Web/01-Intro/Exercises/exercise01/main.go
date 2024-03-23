package main

import (
	"fmt"
	"net/http"
)

/*
Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación web con el package net/http nativo de go, que tenga un endpoint /ping
que al pegarle responda un texto que diga “pong”

El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error", err)
	} else {
		fmt.Println("Server running on port 8080")
	}
}
