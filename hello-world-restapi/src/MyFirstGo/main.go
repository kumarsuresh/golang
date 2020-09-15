package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int
	Body interface{}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("Request get method received")
			data := Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			// data :=  map[string]float64{
			// 	"AMN": 234.44,
			// 	"GOOG": 783.33,
			// 	"TSL":263.14,
			// }
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":8001", mux)
}
