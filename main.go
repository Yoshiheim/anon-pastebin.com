package main

import (
	"go-virtual-currency/controller"
	"go-virtual-currency/db"
	"log"
	"net/http"
	"os"
)

func main() {
	db.Connect()
	controller.InitControllers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	//http.Handle("/mid-test", middle.LoggingMiddleware(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {})))

	// Раздача фронтенда
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fs)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at http://localhost:" + port)
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))
	if err := http.ListenAndServe("localhost:"+port, nil); err != nil {
		panic(err)
	}

}
