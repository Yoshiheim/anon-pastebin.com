package controller

import (
	"go-virtual-currency/handlers"
	"go-virtual-currency/helpers"
	"net/http"
)

func InitPastes() {
	//only for linux
	http.Handle("/4linux/pastes", helpers.IsLinux(http.HandlerFunc(handlers.RenderPastesWithHtml)))

	http.Handle("/create-paste", helpers.LimitMiddleware(http.HandlerFunc(handlers.CreatePaste)))
	http.Handle("/pastes/del", helpers.LimitMiddleware(http.HandlerFunc(handlers.DeletePaste)))

	//for everyone!
	http.HandleFunc("/", handlers.RenderPastesWithHtml)
	http.HandleFunc("/pastes", handlers.RenderPastes)
	http.HandleFunc("/pastes/local", handlers.RenderLocalPaste)
	http.HandleFunc("/pastes/html/view", handlers.ViewPaste)
}
