package _interface

import (
	"log"
	"net/http"
	"testing"
)

func TestDefaultMuxHandler(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
