package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func MakeHandlers() *http.ServeMux {
	mux := *http.NewServeMux()
	mux.HandleFunc("/", sendHello)
	return &mux
}

func sendError(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json")

	status := Response{
		Message: err.Error(),
	}
	fmt.Print(status.Message)
	statusJson, err := json.Marshal(status)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).Write(statusJson)
}

func sendHello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	resp := &Response{
		Message: "Hello",
	}

	byteResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(byteResp)
}
