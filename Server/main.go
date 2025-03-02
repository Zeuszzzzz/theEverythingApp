package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8000
	fmt.Println("Listening at http://127.0.0.1:8000")
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(fmt.Sprint("127.0.0.1:", port), nil)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Hello World!")
}

func homePage(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Add("Context-type", "text/html")
	res.Write([]byte("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Document</title></head><body><p>Hello World!</p><hr /></body></html>"))
	return
}
