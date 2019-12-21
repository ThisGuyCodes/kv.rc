package kv

import (
	"encoding/json"
	"log"
	"net/http"
)

// Getter represents our read into a datastore
type Getter interface {
	// Get has the same interface as a map
	Get(string) (string, bool)
}

// GetHandler provides a net/http interface to a datastore
type GetHandler struct {
	Getter
}

func (gh *GetHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("Bad form parse on get: %q", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	reqKey, ok := req.Form["key"]

	if !ok {
		log.Println("No key requested on get")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(reqKey) > 1 {
		log.Println("Multiple keys requested on get")
		rw.WriteHeader(http.StatusNotImplemented)
		return
	}

	value, ok := gh.Get(reqKey[0])
	if !ok {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	jw := json.NewEncoder(rw)
	err = jw.Encode(map[string]string{reqKey[0]: value})
	if err != nil {
		log.Printf("Error writing result on get: %q", err)
	}
}
