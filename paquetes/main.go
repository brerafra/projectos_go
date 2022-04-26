package main

import (
	"fmt"

	despedida "github.com/brerafra/projectos_go/paquetes/despedida"
	"github.com/brerafra/projectos_go/paquetes/saludar"
)

func main() {
	saludar.Saludar("Brenthon")
	saludar.MeVes = "Esto puede ser un string asignado desde el main a la variable en el paquete"
	fmt.Println(saludar.MeVes)
	
	nombre := "arley"
	despedida.Despedirse(nombre)
}