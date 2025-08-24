package controller

import (
	"go-virtual-currency/handlers"
	"go-virtual-currency/helpers"
	"net/http"
)

func InitPastes() {
	//only for linux
	http.Handle("/4linux/pastes", helpers.IsLinux(http.HandlerFunc(handlers.RenderPastesWithHtml)))

	//for everyone!
	http.HandleFunc("/", handlers.RenderPastesWithHtml)
	http.HandleFunc("/pastes", handlers.RenderPastes)
	http.HandleFunc("/pastes/local", handlers.RenderLocalPaste)
	http.HandleFunc("/pastes/html/view", handlers.ViewPaste)
	http.HandleFunc("/pastes/del", handlers.DeletePaste)
	http.HandleFunc("/create-paste", handlers.CreatePaste)
	http.HandleFunc("/pastes/requests", handlers.GetReq)
	http.HandleFunc("/pastes/2posts", handlers.Create2Posts)
}
