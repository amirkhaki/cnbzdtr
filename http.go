package main

import (
	"net/http"
)

func alive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

