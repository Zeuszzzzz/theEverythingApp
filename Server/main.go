package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	file := os.DirFS("../Frontend/dist")
	fs := http.FS(file)
	handler := http.FileServer(fs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		endPointLogging(r)

		handler.ServeHTTP(w, r)
	})
	http.HandleFunc("/Login", func(w http.ResponseWriter, r *http.Request) {
		endPointLogging(r)

		if r.Method == http.MethodPost {

			fmt.Print("hello world!\n", r.FormValue("EMAIL"), "\n", r.FormValue("PASSWORD"))
		}

		http.ServeFileFS(w, r, file, "index.html")
	})

	fmt.Printf("http://127.0.0.1:%d\n", port)
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
}

func endPointLogging(r *http.Request) {
	res, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(res))
}

func Login() {

}
