package tests

import (
	"testing"

	"github.com/mydearniko/heleket-go"

	"github.com/stretchr/testify/require"
)

func TestCreateStaticWallet(t *testing.T) {
	staticWalletReq := &heleket.StaticWalletRequest{
		Currency: "TRX",
		Network:  "tron",
		OrderId:  "xxx",
		StaticWalletRequestOptions: &heleket.StaticWalletRequestOptions{
			UrlCallback: "https://example.com/heleket/callback",
		},
	}

	staticWallet, err := TestHeleket.CreateStaticWallet(staticWalletReq)
	require.NoError(t, err)
	require.NotEmpty(t, staticWallet)
}

func TestBlockAddressValidation(t *testing.T) {
	// Test that at least one identifier is required
	_, err := TestHeleket.BlockAddress(&heleket.BlockAddressRequest{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "you must provide one of")
}
