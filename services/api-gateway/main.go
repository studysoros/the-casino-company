package main

import (
	"log"
	"net/http"

	"github.com/studysoros/the-casino-company/services/api-gateway/grpc_clients"
	"github.com/studysoros/the-casino-company/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	cashierWrapper, err := grpc_clients.NewCashierServiceClient()
	if err != nil {
		log.Fatalf("Failed to create cashier client: %v", err)
	}
	defer cashierWrapper.Close()

	balanceWrapper, err := grpc_clients.NewBalanceServiceClient()
	if err != nil {
		log.Fatalf("Failed to create balance client: %v", err)
	}
	defer balanceWrapper.Close()

	bettingWrapper, err := grpc_clients.NewBettingServiceClient()
	if err != nil {
		log.Fatalf("Failed to create betting client: %v", err)
	}
	defer bettingWrapper.Close()

	apiServer := &Server{
		CashierClient: cashierWrapper.Client,
		BalanceClient: balanceWrapper.Client,
		BettingClient: bettingWrapper.Client,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /cashier/deposit", apiServer.handleDeposit)
	mux.HandleFunc("POST /cashier/withdraw", apiServer.handleWithdraw)
	mux.HandleFunc("/balance/balance", apiServer.handleGetBalance)
	mux.HandleFunc("POST /betting/placebet", apiServer.handlePlaceBet)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
