package main

import (
	"log"
	"net/http"
)

func main() {
	//Servidor web de forma tradicional

	/*
		//Ruta a neustro servidor que va a ser  la raiz "/" y que debe hacer al realizar esa peticion
		http.Handler("/",http.FileServer(http.Dir("public")))
		//se carga el servidor diciendole que se sirva en un puerto
		//Si el puerto esta ocupado va a devolver un error y para capturar ese error lo hacemos con un log
		log.Println(http.ListenAndServe(":8080",nil))
	*/

	//Servidor web mejorado con un multiplexor (ayuda a resolver las solicitudes)
	//Server mux rsuelve mas peticiones
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/",fs)
	log.Println("ejecutado server en http://localhost:8088")
	log.Println(http.ListenAndServe(":8088",mux))

}