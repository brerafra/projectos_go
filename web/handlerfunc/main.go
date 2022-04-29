package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"<h1>Hola este es un handler func</h1>")
}

//Tenemos una funcion que recibe un mensaje y devuelve un handler
//hander func personalizado

func messageHFCustom(message string)http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, message)
		},
	)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",messageHandlerFunc)

	//personalizando(custom)

	mux.Handle("/saludar", messageHFCustom("<h1> Hola a todos</h1>"))
	mux.Handle("/despedir", messageHFCustom("<h1> Chao a todos</h1>"))

	log.Println("Ejecutando server en http://localhost:8088")
	log.Println(http.ListenAndServe(":8088",mux))
}