package main

import "fmt"

func main() {
	fmt.Println("webos putos")
	//primer manera de definir variables
	var nombre, apellido string
	nombre = "Brenthon"
	apellido = "Ramirez"

	//Segunda manera de definir variables
	apellido2:="Franco"
	//Tercera manera de definir variables

	usuario1, usuario2 := "Karen", "mati"

	fmt.Println(nombre, apellido, apellido2)
	fmt.Println(usuario1, usuario2)
}
