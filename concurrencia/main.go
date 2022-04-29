package main

import (
	"fmt"
	"time"
)

func main() {
	var h string
	go MostrarNumeros()
	go MostrarNombre()
	for{
		fmt.Println("Digita algo: ")
		fmt.Scan(&h)
		fmt.Println("Digitaste: ",h)
	}
}

func MostrarNumeros(){
	for i:=1;i<=1000;i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func MostrarNombre(){
	for {
		fmt.Println("Brenthon Brenthon")
		time.Sleep(time.Second)
	}
}