package main

import (
	"encoding/json"
	"net/http"
)

// Explanation: https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk3
func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

type request struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type response struct {
	Result int `json:"result"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req *request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Unable to decode data", http.StatusInternalServerError)
		return
	}

	var resp response

	if req.Y == 0 {
		http.Error(w, "Unable to divide by 0", http.StatusBadRequest)
		return
	}

	resp.Result = req.X / req.Y
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Unable to encode data", http.StatusInternalServerError)
		return
	}
}
