package main

import (
	"encoding/json"
	"net/http"

	"github.com/studysoros/the-casino-company/shared/contracts"
)

func handleDeposit(w http.ResponseWriter, r *http.Request) {
	var reqBody depositRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: "ok"}
	writeJSON(w, http.StatusCreated, response)
}
