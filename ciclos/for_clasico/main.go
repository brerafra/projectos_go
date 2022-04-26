package main

import "fmt"

func main() {
	//Ejemplo 1
	/*
		count:=5
		for i:=0;i<count;i++{
			fmt.Printf(i)
		}
	*/

	//Ejemplo 2
	/*
		for i:=5; i>=0; i--{
			if i==3{
				fmt.Println("Se salta el valor de 3")
				continue
			}
			fmt.Println(i)
		}
	*/

	//ejemplo3
	/*
		count:=5
		for i:=0;i<count;i++{
			if i==2{
				fmt.Println("Cuando i vale 2 se rompe el ciclo")
				break
			}
			fmt.Println(i)
		}
	*/

	//Ejemplo 4
	matriz := [3][4]int{}
	valor := 0
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			valor++
			matriz[i][y] = valor
		}
	}
	for i := 0; i < 3; i++ {
		for y := 0; y < 3; y++ {
			fmt.Println(matriz[i][y])
		}
	}
}