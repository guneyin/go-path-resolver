package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = 3001

func handlePost(data []byte) []byte {
	var r Request
	_ = json.Unmarshal(data, &r)

	r.Init()

	b, err := json.Marshal(r.Response)
	if err != nil {
		r := fmt.Sprintf(`{"error": "%v"}`, err.Error())
		return []byte(r)
	}

	return b
}

func apiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		r := handlePost(body)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(r)

	default:
		r := fmt.Sprintf(`{"error": "method %v not allowed"}`, r.Method)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(r))
	}
}

func main() {
	http.HandleFunc("/", apiResponse)

	log.Println(fmt.Sprintf("Listening on port %v", PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}
