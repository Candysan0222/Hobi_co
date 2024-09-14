package login

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := "Pepe registrado"
	json.NewEncoder(w).Encode(p)
}
