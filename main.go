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

		fmt.Println(i)
	}
}
func ordenarArray() {
	a := []string{"Cristina", "Rafael", "Leonardo", "Yanina", "Emmanuel", "Aaron"}
	sort.Strings(a)
	fmt.Println(a)
}
func nombreCompleto(nombre string, apellido string) {
	nombre = "Cristina"
	apellido = " Ahumada"

	nombreCompleto("Bienvenida ", nombre+apellido)

}
