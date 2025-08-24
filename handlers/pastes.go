package handlers

import (
	"encoding/json"
	"fmt"
	"go-virtual-currency/db"
	"go-virtual-currency/helpers"
	"go-virtual-currency/models"
	"html"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type ModelJson struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// рендерим все пасты
func RenderPastes(w http.ResponseWriter, r *http.Request) {

	var pastes []models.Paste

	gets := db.DB.Find(&pastes)
	if gets.Error != nil {
		log.Println(gets.Error)
		http.Error(w, "	Cannot Create...", http.StatusNotAcceptable)
		return
	}

	log.Printf("pastes: %v\n", gets)

	helpers.EncodeJson(w, map[string]interface{}{
		"data": pastes,
	})
}

type IdRequrest struct {
	ID uint `json:"id"`
}

func RenderLocalPaste(w http.ResponseWriter, r *http.Request) {

	var body IdRequrest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Println(err)
		http.Error(w, "Error with decode", http.StatusBadRequest)
		return
	}

	var pastes models.Paste

	gets := db.DB.Find(&pastes).Where("ID = ?", body.ID)
	if gets.Error != nil {
		log.Println(gets.Error)
		http.Error(w, "	Cannot Create...", http.StatusNotAcceptable)
		return
	}

	log.Printf("pastes: %v\n", pastes.Title)

	w.Header().Set("Content-Type", "application/json")

	helpers.EncodeJson(w, map[string]interface{}{
		"id":      pastes.ID,
		"title":   pastes.Title,
		"content": pastes.Content,
	})
	log.Printf("%d - %s - %s\n", pastes.ID, pastes.Title, pastes.Content)
}

func RenderPastesWithHtml(w http.ResponseWriter, r *http.Request) {
	var pastes []models.Paste

	gets := db.DB.Find(&pastes)
	if gets.Error != nil {
		log.Println(gets.Error)
		http.Error(w, "	Cannot Create...", http.StatusNotAcceptable)
		return
	}

	for i, _ := range pastes {
		pastes[i].Content = html.EscapeString(pastes[i].Content)
		pastes[i].Title = html.EscapeString(pastes[i].Title)
	}

	templ, err := template.ParseFiles("./templates/pastes.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	log.Println(r.UserAgent())
	templ.Execute(w, pastes)
}

func ViewPaste(w http.ResponseWriter, r *http.Request) {
	var paste models.Paste

	id := r.URL.Query().Get("id")

	idnum, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Wront token %s\n", id)
		http.Error(w, fmt.Sprintf("Wront token %s\n", id), http.StatusBadRequest)
		return
	}

	get := db.DB.Where("ID = ?", idnum).Find(&paste)
	if get.Error != nil {
		log.Println(get.Error.Error())
		http.Error(w, get.Error.Error(), http.StatusNotAcceptable)
		return
	}
	paste.Content = html.EscapeString(paste.Content)
	paste.Title = html.EscapeString(paste.Title)
	templ, err := template.ParseFiles("./templates/localpaste.html")
	if err != nil {
		log.Printf("i cant parse file: %s\n", err)
		http.Error(w, fmt.Sprintf("i cant parse file: %s\n", err), http.StatusNotAcceptable)
		return
	}
	if err := templ.Execute(w, paste); err != nil {
		log.Printf("i cant parse file: %s\n", err)
		http.Error(w, fmt.Sprintf("i cant parse file: %s\n", err), http.StatusNotAcceptable)
		return
	}
}

func CreatePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var mj ModelJson

	// Чтение JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&mj); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}
	if mj.Title == "" || mj.Content == "" {
		log.Println("body is empty")
		http.Error(w, "body is empty", http.StatusBadRequest)
		return
	}

	// Сохраняем в БД
	cre := db.DB.Create(&models.Paste{
		Title:   mj.Title,
		Content: mj.Content,
	})

	if cre.Error != nil {
		log.Println("DB Create Error:", cre.Error)
		http.Error(w, "Cannot create paste", http.StatusNotAcceptable)
		return
	}

	// Ответ (по желанию)
	w.WriteHeader(http.StatusCreated)
}

func GetReq(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	site, err := http.Get("https://e621.net")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error with get...", http.StatusNotAcceptable)
		return
	}
	buff, err := io.ReadAll(site.Body)
	if err != nil {
		http.Error(w, "Error with copy", http.StatusBadGateway)
		return
	}
	cre := db.DB.Create(&models.Paste{
		Title:   "test",
		Content: string(buff),
	})

	if cre.Error != nil {
		log.Println("DB Create Error:", cre.Error)
		http.Error(w, "Cannot create paste", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeletePaste(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Canot convert id", http.StatusBadRequest)
		return
	}
	del := db.DB.Where("ID = ?", num).Delete(&models.Paste{})
	if del.Error != nil {
		log.Println(err)
		http.Error(w, "Error with db", http.StatusNotAcceptable)
		return
	}
	http.Error(w, "okay", http.StatusOK)
}

type URLBro struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

func Create2Posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var ub URLBro

	if err := json.NewDecoder(r.Body).Decode(&ub); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Cannot parse JSON body", http.StatusBadRequest)
		return
	}

	site, err := http.Get(ub.First)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error with get...", http.StatusNotAcceptable)
		return
	}
	buff, err := io.ReadAll(site.Body)
	if err != nil {
		http.Error(w, "Error with copy", http.StatusBadGateway)
		return
	}
	cre := db.DB.Create(&models.Paste{
		Title:   "first",
		Content: string(buff),
	})

	if cre.Error != nil {
		log.Println("DB Create Error:", cre.Error)
		http.Error(w, "Cannot create paste", http.StatusNotAcceptable)
		return
	}

	site2, err := http.Get(ub.Second)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error with get...(second)", http.StatusNotAcceptable)
		return
	}
	buff2, err := io.ReadAll(site2.Body)
	if err != nil {
		http.Error(w, "Error with copy(second)", http.StatusBadGateway)
		return
	}

	cre2 := db.DB.Create(&models.Paste{
		Title:   "second",
		Content: string(buff2),
	})

	if cre2.Error != nil {
		log.Println("DB Create Error(second):", cre2.Error)
		http.Error(w, "Cannot create paste(second)", http.StatusNotAcceptable)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
