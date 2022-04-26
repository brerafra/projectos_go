package main

import "fmt"

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"}

	//El range devuelve dos valores el indice y el valor
	for i, v := range numeros {
		fmt.Println(i, v)
	}

	//si no queremos utilizar el indice coloca _
	for _, v := range numeros {
		fmt.Println(v)
	}

	//El valor se puede omitir en un range
	for i:=range numeros {
		fmt.Println(i)
	}

	//ejemplo con mapas funciona de la misma manera
	nombres:=map[string]string{
		"es":"España",
		"co":"Colombia",
		"br":"Brazil",
	}
	for i,v := range nombres {fmt.Println(i,v)}

	//iteracion con string
	frase := "Hola mundo cruel"
	for posicion,letra:= range frase{
		fmt.Println(posicion,string(letra))
	}

	//iterar slices de enteros
	for _,entero:=range[]int{15,36,24,86,15,23,187}{
		fmt.Println(entero)
	}

	//Iterar un string
	for _,letra:= range "Hola Mundo"{
		fmt.Println(string(letra))
	}
}