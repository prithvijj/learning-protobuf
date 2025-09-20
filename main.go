package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"learning-protobuf-go/examplepb"

	"google.golang.org/protobuf/proto"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var p examplepb.Person
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(&p)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func protoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var p examplepb.Person
	if err := proto.Unmarshal(body, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := proto.Marshal(&p)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(resp)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/proto", protoHandler)

	log.Println("Server listening on :8765")
	log.Fatal(http.ListenAndServe(":8765", nil))
}
