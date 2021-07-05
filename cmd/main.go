package main

import (
	"contra-design.com/govue/cmd/db"
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplIndex *template.Template
)

type Error struct {
	Message string `json:"message"`
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var dir string

	dsn := os.Getenv("DATABASE_URL")
	db.DBInit(dsn)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplIndex = template.Must(template.ParseFiles("index.gohtml"))

	r := mux.NewRouter()
	flag.StringVar(&dir, "dir", "dist", "the directory to serve files from")
	flag.Parse()

	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/api/register", apiRegister).Methods("POST")
	r.HandleFunc("/api/login", apiLogin).Methods("POST")
	r.HandleFunc("/api/test", TokenVerifyMiddleWare(apiTest)).Methods("GET")

	r.HandleFunc("/{category}/{name}/{id}", index)
	r.HandleFunc("/{category}/{name}", index)
	r.HandleFunc("/{name}", index)
	r.HandleFunc("/", index).Methods("GET")

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":" + port, r)
}
