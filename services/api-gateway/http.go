package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/studysoros/the-casino-company/shared/contracts"
	pbBalance "github.com/studysoros/the-casino-company/shared/proto/balance"
	pbBetting "github.com/studysoros/the-casino-company/shared/proto/betting"
	pbCashier "github.com/studysoros/the-casino-company/shared/proto/cashier"
)

type Server struct {
	BalanceClient pbBalance.BalanceServiceClient
	BettingClient pbBetting.BettingServiceClient
	CashierClient pbCashier.CashierServiceClient
}

func (s *Server) handleDeposit(w http.ResponseWriter, r *http.Request) {
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

	deposit, err := s.CashierClient.Deposit(ctx, reqBody.toProto())
	if err != nil {
		log.Printf("Failed to deposit: %v", err)
		http.Error(w, "Failed to deposit", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: deposit}
	writeJSON(w, http.StatusCreated, response)
}

func (s *Server) handleWithdraw(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody withdrawRequest
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

	withdraw, err := s.CashierClient.Withdraw(ctx, reqBody.toProto())
	if err != nil {
		log.Printf("Failed to withdraw: %v", err)
		http.Error(w, "Failed to withdraw", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: withdraw}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handleGetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody getBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	balance, err := s.BalanceClient.GetBalance(ctx, reqBody.toProto())
	if err != nil {
		log.Printf("Failed to get user balance: %v", err)
		http.Error(w, "Failed to get user balance", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: balance}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) handlePlaceBet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody betRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data.", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if reqBody.BetSide == "" {
		http.Error(w, "Please place your side", http.StatusBadRequest)
		return
	}

	bet, err := s.BettingClient.PlaceBet(ctx, reqBody.toProto())
	if err != nil {
		log.Printf("Failed to place a bet: %v", err)
		http.Error(w, "Failed to place a bet", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: bet}
	writeJSON(w, http.StatusOK, response)
}
