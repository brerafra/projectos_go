package main

import (
	"fmt"

	"github.com/dickson7/go/13_paquetes/saludar"
)

func main() {
	saludar.Saludar("Brenthon")
	saludar.MeVes = "Esto puede ser un string asignado desde el main a la variable en el paquete"
	fmt.Println(saludar.MeVes)
	
	nombre := "arley"
	depedida.Despedirse(nombre)
}