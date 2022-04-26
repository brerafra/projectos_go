package main

import "fmt"

func main() {
	var n1, n2 int8
	n1 = 10
	n2 = 17

	r := suma(n1, n2)
	fmt.Println(r)

	var edad uint8
	edad=3
	fmt.Println(tipoEdad(edad))

	var n[]uint8
	n=[]uint8{8,2,6,7,9,1}
	fmt.Println(n)
	maximo,minimo :=maxymin(n)
	fmt.Printf("El valor mayor es %d y el valor menos %d \n",maximo,minimo)
}

func suma(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "niñez"
	case edad < 18:
		tipo = "adolescencia"
	default:
		tipo = "madurez"
	}
	return tipo
}

//manera de firmar las funciones
/*
	func maxymin(numeros []uint8)(uint8, uint8) {

	}
*/

func maxymin(numeros []uint8) (maxi uint8, mini uint8) {
	for _, v := range numeros {
		if v > maxi {
			maxi = v
		}
		if v < mini {
			mini = v
		}
	}
	return
}