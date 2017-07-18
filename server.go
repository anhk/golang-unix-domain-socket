package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Handler: mux, // http.DefaultServeMux,
	}

	os.Remove("/var/run/test.sock")

	listener, err := net.Listen("unix", "/var/run/test.sock")
	if err != nil {
		panic(err)
	}
	server.Serve(listener)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQ: ", r.URL)
	fmt.Fprintf(w, "Hello World!")
}
