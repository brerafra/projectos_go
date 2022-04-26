package main

import "fmt"

func main() {
	//En los arrays todos los valores son de
	//tamaño fijo y del mismo tipo de datos

	var nombres [3]string
	nombres[0] = "Brenthon"
	nombres[1] = "Ramirez"
	nombres[2] = "Franco"

	fmt.Println(nombres)

	apellidos :=[3]string{"Garcia","Lopez","perez"}
	fmt.Println(apellidos)
	fmt.Println(nombres[1],apellidos[1])

	//metodo del arreglo para conocer su tamaño 
	size:=len(nombres)
	fmt.Println("El tamaño del arreglo es de: ",size)

	//fmt.Println(nombres[inicio:fina(excluyente)])
	fmt.Println(nombres[0:2])
}