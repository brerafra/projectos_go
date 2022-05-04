package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID int `json:id`
	Name  string `json:Name`
	Content string `json:Content`
}

func getTasks(w http.ResponseWriter, r *http.Request){
	
}

type allTask []task

var tasks = allTask{
	{
		ID:1,
		Name:"Task One",
		Content:"Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request){
	dato1:= 25
	dato2:= 55

	suma:= dato1+dato2
	fmt.Fprintf(w,"Bienvenido a mi API: ",suma)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",indexRoute)

	log.Fatal(http.ListenAndServe(":8081",router))
}