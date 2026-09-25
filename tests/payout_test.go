package tests

import (
	"testing"

	"github.com/mydearniko/heleket-go"
	"github.com/stretchr/testify/require"
)

func TestGetPayoutServicesList(t *testing.T) {
	services, err := TestHeleket.GetPayoutServicesList()
	require.NoError(t, err)
	require.NotNil(t, services)
	// Services list might be empty, but should not error
}

func TestGetPayoutInfoValidation(t *testing.T) {
	// Test that at least one identifier is required
	_, err := TestHeleket.GetPayoutInfo(&heleket.PayoutInfoRequest{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "you must provide one of")
}
