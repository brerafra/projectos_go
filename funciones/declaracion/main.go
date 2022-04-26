package main

import "fmt"

func main() {
	saludar("Brenthon", 33)
}

func saludar(name string, edad uint8) {
	fmt.Println("Hola ", name)
	fmt.Printf("Tienes %d años \n",edad)
	if edad > 30{
		fmt.Println("Estas viejo jaja")
		return 
	}
	fmt.Println("Eres joven")
}