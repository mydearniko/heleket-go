package tests

import (
	"testing"

	"github.com/mydearniko/heleket-go"
	"github.com/stretchr/testify/require"
)

func TestRefundValidation(t *testing.T) {
	// Test that at least one identifier is required
	_, err := TestHeleket.Refund(&heleket.RefundRequest{
		Address:    "test-address",
		IsSubtract: true,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "you must provide one of")
}

func TestBlockedAddressRefundValidation(t *testing.T) {
	// Test that at least one identifier is required
	_, err := TestHeleket.BlockedAddressRefund(&heleket.BlockedAddressRefundRequest{
		Address: "test-address",
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "you must provide one of")
}
