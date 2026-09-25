package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBalance(t *testing.T) {
	balance, err := TestHeleket.GetBalance()
	require.NoError(t, err)
	require.NotNil(t, balance)
	// Balance should never be nil, even if empty
	require.NotNil(t, balance.Merchant)
	require.NotNil(t, balance.User)
}

func TestGetDiscountsList(t *testing.T) {
	discounts, err := TestHeleket.GetDiscountsList()
	require.NoError(t, err)
	require.NotNil(t, discounts)
}

func TestGetExchangeRates(t *testing.T) {
	rates, err := TestHeleket.GetExchangeRates("BTC")
	require.NoError(t, err)
	require.NotNil(t, rates)
}

func TestGetPaymentServicesList(t *testing.T) {
	services, err := TestHeleket.GetPaymentServicesList()
	require.NoError(t, err)
	require.NotNil(t, services)
}
