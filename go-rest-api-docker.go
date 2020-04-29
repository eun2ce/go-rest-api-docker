package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type data struct {
	code        string
	Title       string
	Description string
}

func GetData(w http.ResponseWriter, r *http.Request) {
	var testData []data
	testData = append(testData, data{code: "1", Title: "first title", Description: "desc"})
	testData = append(testData, data{code: "2", Title: "second title", Description: "desc"})

	p := mux.Vars(r)
	for _, i := range testData {
		if i.code == p["code"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&data{})
}

//2. drive server
func httpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// 1. Create router
	router := mux.NewRouter()
	//3. register router func
	///register getData func on /test/code
	router.HandleFunc("/test/{code}", GetData).Methods("GET")

	// 1. Create router
	http.ListenAndServe(":8080", router)

}
