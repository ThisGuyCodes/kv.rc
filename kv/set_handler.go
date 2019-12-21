package kv

import (
	"log"
	"net/http"
)

// Setter represents our write into a datastore
type Setter interface {
	Set(string, string)
}

// SetHandler provides a net/http interface to a datastore
type SetHandler struct {
	Setter
}

func (sh *SetHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("Bad form parse on set: %q", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.Form) > 1 {
		log.Println("Attempt to set multiple keys on set")
		rw.WriteHeader(http.StatusNotImplemented)
		return
	}

	if len(req.Form) == 0 {
		log.Println("No key specified on set")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// this will only run-through once
	for reqKey, values := range req.Form {
		if len(values) > 1 {
			log.Println("Attempt to set same key multiple times on set")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		sh.Set(reqKey, values[0])
	}
}
