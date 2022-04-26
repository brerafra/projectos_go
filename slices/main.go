package main

import "fmt"

func main() {
	//Slices tienen un apuntador a un array
	//tienen un tamaño y tambien una capacidad

	//Declaracion de slices
	// var nombres[]string

	//otra manera de declarar slices
	nombres := make([]string, 0)
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"Brenthon")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"KAren")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"Mati")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"Juan")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"Astrid")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))
	nombres =append(nombres,"Ale")
	fmt.Printf("Su tamaño es: %d y su capacidad es: %d\n",len(nombres), cap(nombres))

	//otra manera de agregar informacion a nuestro slicen
	apellidos:=[]string{
		"Garcia",
		"Lopez",
		"Perez",
	}
	fmt.Println(nombres)
	fmt.Println(apellidos)
}