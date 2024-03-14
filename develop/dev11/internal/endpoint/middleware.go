package endpoint

import (
	"log"
	"net/http"
)

func middleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request:\nMethod: %s\nURI: %s\nBody: %s\n", r.Method, r.RequestURI, r.Body)
		hf(w, r)
		log.Printf("Response: %v\n\n", w)
	}
}