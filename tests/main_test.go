package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/idanyas/heleket-go"
)

var TestHeleket *heleket.Heleket

func TestMain(m *testing.M) {
	httpClient := &http.Client{}
	
	// Use environment variables for credentials
	merchant := os.Getenv("HELEKET_MERCHANT_ID")
	paymentAPIKey := os.Getenv("HELEKET_PAYMENT_API_KEY")
	payoutAPIKey := os.Getenv("HELEKET_PAYOUT_API_KEY")
	
	// Skip tests if credentials are not provided
	if merchant == "" || paymentAPIKey == "" || payoutAPIKey == "" {
		println("Skipping tests: HELEKET_MERCHANT_ID, HELEKET_PAYMENT_API_KEY, and HELEKET_PAYOUT_API_KEY environment variables must be set")
		os.Exit(0)
	}

	client, err := heleket.New(httpClient, merchant, paymentAPIKey, payoutAPIKey)
	if err != nil {
		println("Failed to create Heleket client:", err.Error())
		os.Exit(1)
	}
	
	TestHeleket = client

	os.Exit(m.Run())
}
