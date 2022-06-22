package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, world!")
	loop()
	ordenarArray()
	nombreCompleto()
}

func loop() {
	for i := 0; i < 60; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Este es el loop", i)
	}
}
func ordenarArray() {
	a := []string{"Cristina", "Rafael", "Leonardo", "Yanina", "Emmanuel", "Aaron"}
	sort.Strings(a)
	fmt.Println("El orden alfabetico de los nombres en el array es", a)
}
func nombreCompleto() {
	var nombre = "Cristina"
	var apellido = " Ahumada"
	var nombreC = nombre + apellido
	fmt.Println("El nombre completo es", nombreC)
}
