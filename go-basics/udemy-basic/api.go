package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("Hi")
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request receive")
		w.Write([]byte("get going is working"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request receive")
		w.Write([]byte("get going is working go ingin gin"))
	})
	http.ListenAndServe("localhost:8001", mux)
}