package main

import "fmt"

func main() {
	//Declaracion de maps
	/*
		idiomas:=make(map[string]string)
		idiomas["es"]="Español"
		idiomas["en"]="Ingles"
		idiomas["fr"]="Frances"
	*/

	//Otra forma de declaracion de maps
	idiomas := map[string]string{
		"es": "Español",
		"en": "Ingles",
		"fr": "Frances",
	}
	//imprimir todo el mapa
	fmt.Println(idiomas)
	//Imprimir solo una posicion en especifico
	fmt.Println(idiomas["en"])
	//eliminar una posicion en el mapa
	delete(idiomas, "en")
	fmt.Println(idiomas)

	//comprobar si existe una posicion
	if idioma, ok :=idiomas["en"];ok{
		fmt.Println("La posicion existe y vale ",idioma)
	}else{
		fmt.Println("LA posicion no existe")
	}

	//declaracion con int y string 
	numeros:=map[int]string{
		1:"un numero pequeño",
		20:"otro numero",
		-3:"llave negativa",
	}
	fmt.Println(numeros)
}