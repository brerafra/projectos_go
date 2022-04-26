package main

import "fmt"

func main() {
	saludarVarios(10, "Brenthon", "karen", "mati")
}

//las funciones variaticas solo reciben un unico parametro variatico
func saludarVarios(edad uint8, nombres ...string) {
	//para saber que tipo de datos es una variable
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "edad", edad)

	}
}