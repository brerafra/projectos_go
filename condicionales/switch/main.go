package main

import "fmt"

func main() {
	//ejemplo 1
	var a uint8
	a = 3
	switch a {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Lunes")
	case 3:
		fmt.Println("Martes")
	case 4:
		fmt.Println("Miercoles")
	case 5:
		fmt.Println("Jueves")
	case 6:
		fmt.Println("Viernes")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("No es un dia de la semana")
	}

	switch a{
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("estas entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Ests en fin de semana")
	default:
		fmt.Println("No es un dia valido")
	}

	switch b:= 7;{
	case b>=0 && b<=5:
		fmt.Println("Estas entre semana")
	case b>=6 && b<=7:
		fmt.Println("Estas en fin de semana")
	default:
		fmt.Println("No es un dia valido")
	}
}