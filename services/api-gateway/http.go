package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/studysoros/the-casino-company/services/api-gateway/grpc_clients"
	"github.com/studysoros/the-casino-company/shared/contracts"
)

func handleDeposit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	if reqBody.Amount == 0 {
		http.Error(w, "Amount cannot be zero", http.StatusBadRequest)
		return
	}

	cashierService, err := grpc_clients.NewCashierServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer cashierService.Close()

	receipt, err := cashierService.Client.Deposit(ctx, reqBody.toProto())
	if err != nil {
		log.Printf("Failed to deposit: %v", err)
		http.Error(w, "Failed to deposit", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: receipt}
	writeJSON(w, http.StatusCreated, response)
}
