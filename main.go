package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	loop()
}

func loop() {
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}
