package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	rtr  *mux.Router
	tmpl *template.Template
)

func main() {

	rtr = mux.NewRouter()
	tmpl = template.Must(template.ParseGlob("dist/*.html"))
	rtr.HandleFunc("/", index)
	rtr.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("dist/assets"))))
	rtr.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("unreal/"))))

	log.Println("Server running at http://localhost:8080")

	//****************** MUST BE END OF MAIN ******************\\
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handlers.LoggingHandler(os.Stdout, rtr),
		MaxHeaderBytes: 1 << 62,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("server failed", err)
	}
	//****************** MUST BE END OF MAIN ******************\\
}

func health(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		"OK",
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(status)
	w.Write(j)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
