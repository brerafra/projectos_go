package main

import "fmt"

func main() {
	//Defer es como una sala de espera
	// que hace la ejecución de codigo cuando se termina la funcion

	fmt.Println("Conectando a una base de datos...")
	defer cerrarDB()

	fmt.Println("Consultamos infomración de set de datos")
	defer cerrarSetDeDatos()
}

func cerrarDB(){
	fmt.Println("Cerrar la BD")
}

func cerrarSetDeDatos(){
	fmt.Println("Cerrar el set de datos")
}