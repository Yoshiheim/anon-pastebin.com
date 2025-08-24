package main

import (
	"go-virtual-currency/controller"
	"go-virtual-currency/db"
	"log"
	"net/http"
)

func main() {

	db.Connect()
	controller.InitControllers()
	//http.Handle("/mid-test", middle.LoggingMiddleware(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {})))

	// Раздача фронтенда
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fs)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
