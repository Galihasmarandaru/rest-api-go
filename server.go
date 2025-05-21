package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Server!")
	})

	// const serverAddress string = "127.0.0.1:3000"
	const port string = ":3000" // bisa gunakan port langsung

	// fmt.Println("Server Listening on port 3000")
	fmt.Println("Server Listening on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting sever", err)
	}
}
