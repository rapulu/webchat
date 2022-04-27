package main

import (
	"fmt"
	"net/http"
	// socketio "github.com/googollee/go-socket.io"
)

func main() {

	// server := socketio.NewServer(nil)
	http.HandleFunc("/", func(resp http.ResponseWriter,req *http.Request) {
		http.ServeFile(resp, req, "view/index.html")
	})


	http.HandleFunc("/style.css", func (resp http.ResponseWriter,req *http.Request) {
		http.ServeFile(resp, req, "view/style.css")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on localhost:8080")
}