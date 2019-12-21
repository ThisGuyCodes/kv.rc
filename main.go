package main

import (
	"github.com/thisguycodes/kv.rc/kv"
	"github.com/thisguycodes/kv.rc/store"
	"net/http"
)

func main() {
	storage := store.New()

	setHandler := &kv.SetHandler{Setter: storage}
	getHandler := &kv.GetHandler{Getter: storage}

	http.Handle("/get", getHandler)
	http.Handle("/set", setHandler)

	http.ListenAndServe(":4000", nil)
}
