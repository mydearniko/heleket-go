package tests

import (
	"net/http"
	"testing"

	"github.com/idanyas/heleket-go"
	"github.com/stretchr/testify/require"
)

func TestNewClientValidation(t *testing.T) {
	tests := []struct {
		name          string
		client        *http.Client
		merchant      string
		paymentKey    string
		payoutKey     string
		expectedError string
	}{
		{
			name:          "nil http client",
			client:        nil,
			merchant:      "test",
			paymentKey:    "test",
			payoutKey:     "test",
			expectedError: "http client cannot be nil",
		},
		{
			name:          "empty merchant",
			client:        &http.Client{},
			merchant:      "",
			paymentKey:    "test",
			payoutKey:     "test",
			expectedError: "merchant ID cannot be empty",
		},
		{
			name:          "empty payment key",
			client:        &http.Client{},
			merchant:      "test",
			paymentKey:    "",
			payoutKey:     "test",
			expectedError: "payment API key cannot be empty",
		},
		{
			name:          "empty payout key",
			client:        &http.Client{},
			merchant:      "test",
			paymentKey:    "test",
			payoutKey:     "",
			expectedError: "payout API key cannot be empty",
		},
		{
			name:          "valid parameters",
			client:        &http.Client{},
			merchant:      "test",
			paymentKey:    "test",
			payoutKey:     "test",
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := heleket.New(tt.client, tt.merchant, tt.paymentKey, tt.payoutKey)
			
			if tt.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError)
				require.Nil(t, client)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})
	}
}
