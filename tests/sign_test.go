package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySignMissingSignField(t *testing.T) {
	payloadWithoutSign := []byte(`{
		"type": "payment",
		"uuid": "test-uuid",
		"amount": "10.00"
	}`)

	err := TestHeleket.VerifySign("test-key", payloadWithoutSign)
	require.Error(t, err)
	require.Contains(t, err.Error(), "missing or invalid 'sign' field")
}

func TestVerifySignInvalidJSON(t *testing.T) {
	invalidJSON := []byte(`{invalid json}`)

	err := TestHeleket.VerifySign("test-key", invalidJSON)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid JSON")
}

func TestVerifySignInvalidSignature(t *testing.T) {
	payloadWithWrongSign := []byte(`{
		"type": "payment",
		"uuid": "test-uuid",
		"amount": "10.00",
		"sign": "wrong-signature"
	}`)

	err := TestHeleket.VerifySign("test-key", payloadWithWrongSign)
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid signature")
}
