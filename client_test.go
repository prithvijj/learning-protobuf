package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"google.golang.org/protobuf/proto"

	"learning-protobuf-go/examplepb"
)

var person = &examplepb.Person{
	Name:  "Alice",
	Age:   30,
	Email: "alice@example.com",
}

func BenchmarkJSONRequest(b *testing.B) {
	data, _ := json.Marshal(person)
	url := "http://localhost:8765/json"

	for i := 0; i < b.N; i++ {
		resp, err := http.Post(url, "application/json", bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}

func BenchmarkProtoRequest(b *testing.B) {
	data, _ := proto.Marshal(person)
	url := "http://localhost:8765/proto"

	for i := 0; i < b.N; i++ {
		resp, err := http.Post(url, "application/x-protobuf", bytes.NewReader(data))
		if err != nil {
			b.Fatal(err)
		}
		resp.Body.Close()
	}
}
