package heleket

import (
	"encoding/json"
	"errors"
)

const (
	refundEndpoint               = "/payment/refund"
	blockedAddressRefundEndpoint = "/wallet/blocked-address-refund"
)

type RefundRequest struct {
	Address     string `json:"address"`
	IsSubtract  bool   `json:"is_subtract"`
	PaymentUUID string `json:"uuid,omitempty"`
	OrderId     string `json:"order_id,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

type refundRawResponse struct {
	Result []string `json:"result,omitempty"`
	State  int8     `json:"state"`
}

type BlockedAddressRefundRequest struct {
	WalletUUID string `json:"uuid,omitempty"`
	OrderId    string `json:"order_id,omitempty"`
	Address    string `json:"address"`
}

type BlockedAddressRefundResponse struct {
	Commission string `json:"commission"`
	Amount     string `json:"amount"`
}

type blockedAddressRefundRawResponse struct {
	Result *BlockedAddressRefundResponse `json:"result"`
	State  int8                          `json:"state"`
}

// Refund processes a refund for a completed payment, sending funds back to a specified address.
// You must provide either PaymentUUID or OrderId in the request.
//
// Parameters:
//   - refundRequest: Contains payment identifier, refund address, and commission handling
//
// Returns true if the refund was successfully initiated.
func (c *Heleket) Refund(refundRequest *RefundRequest) (bool, error) {
	if refundRequest.PaymentUUID == "" && refundRequest.OrderId == "" {
		return false, errors.New("you must provide one of: [PaymentUUID, OrderId]")
	}

	res, err := c.fetch("POST", refundEndpoint, refundRequest, c.paymentApiKey)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	response := &refundRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return false, err
	}

	return len(response.Result) == 0, nil
}

// BlockedAddressRefund processes a refund from a blocked static wallet address.
// You must provide either WalletUUID or OrderId in the request.
//
// Returns details about the refund including amount and commission.
func (c *Heleket) BlockedAddressRefund(refundRequest *BlockedAddressRefundRequest) (*BlockedAddressRefundResponse, error) {
	if refundRequest.WalletUUID == "" && refundRequest.OrderId == "" {
		return nil, errors.New("you must provide one of: [WalletUUID, OrderId]")
	}

	res, err := c.fetch("POST", blockedAddressRefundEndpoint, refundRequest, c.paymentApiKey)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response := &blockedAddressRefundRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
