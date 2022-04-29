package main

import (
	"log"
	"net/http"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   uint8
}

func renderizar(w http.ResponseWriter, r *http.Request) {
	p := &Persona{"Brenthon", 33}
	//El paquete template procesa los archivos y devuelve el template y un error y se controlan
	t, err :=template.ParseFiles("./views/index.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,p)
}


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",renderizar)

	log.Println("Ejecutando server en http://localhost:8088")
	log.Println(http.ListenAndServe(":8088",mux))
	log.Println("Servidor no esta corriendo")
}