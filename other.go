package heleket

import (
	"encoding/json"
	"fmt"
)

const (
	balanceEndpoint             = "/balance"
	discountListEndpoint        = "/payment/discount/list"
	discountSetEndpoint         = "/payment/discount/set"
	exchangeRateListEndpointFmt = "/exchange-rate/%s/list"
)

// Balance
type WalletBalance struct {
	UUID         string `json:"uuid"`
	Balance      string `json:"balance"`
	CurrencyCode string `json:"currency_code"`
}

type BalanceInfo struct {
	Merchant []*WalletBalance `json:"merchant"`
	User     []*WalletBalance `json:"user"`
}

type balanceRawResponse struct {
	State  int8 `json:"state"`
	Result []struct {
		Balance *BalanceInfo `json:"balance"`
	} `json:"result"`
}

// Discount
type Discount struct {
	Network  string `json:"network"`
	Currency string `json:"currency"`
	Discount int8   `json:"discount"`
}

type discountListRawResponse struct {
	State  int8        `json:"state"`
	Result []*Discount `json:"result"`
}

type SetDiscountRequest struct {
	Currency        string `json:"currency"`
	Network         string `json:"network"`
	DiscountPercent int8   `json:"discount_percent"`
}

type setDiscountRawResponse struct {
	State  int8      `json:"state"`
	Result *Discount `json:"result"`
}

// Exchange Rate
type ExchangeRate struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Course string `json:"course"`
}

type exchangeRateListRawResponse struct {
	State  int8            `json:"state"`
	Result []*ExchangeRate `json:"result"`
}

// GetBalance retrieves the current balance for all currencies in your merchant and user accounts.
// Returns BalanceInfo with separate balances for merchant and user wallets.
// If no balances exist, returns an empty BalanceInfo (not nil).
func (c *Heleket) GetBalance() (*BalanceInfo, error) {
	res, err := c.fetch("POST", balanceEndpoint, make(map[string]any), c.paymentApiKey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &balanceRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	if len(response.Result) > 0 && response.Result[0].Balance != nil {
		return response.Result[0].Balance, nil
	}

	// Return empty balance info instead of nil
	return &BalanceInfo{
		Merchant: []*WalletBalance{},
		User:     []*WalletBalance{},
	}, nil
}

// GetDiscountsList retrieves all configured payment method discounts for your merchant account.
// Discounts can be positive (customer discount) or negative (additional commission).
func (c *Heleket) GetDiscountsList() ([]*Discount, error) {
	res, err := c.fetch("POST", discountListEndpoint, make(map[string]any), c.paymentApiKey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &discountListRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// SetDiscount configures a discount or additional commission for a specific payment method.
// Use positive DiscountPercent for customer discounts, negative for additional commission.
//
// Example: DiscountPercent = -5 means add 5% commission for the customer.
func (c *Heleket) SetDiscount(req *SetDiscountRequest) (*Discount, error) {
	res, err := c.fetch("POST", discountSetEndpoint, req, c.paymentApiKey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &setDiscountRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}

// GetExchangeRates retrieves current exchange rates for a specific currency.
// Returns conversion rates from the specified currency to other supported currencies.
//
// Parameters:
//   - currency: The base currency code (e.g., "BTC", "ETH", "USDT")
func (c *Heleket) GetExchangeRates(currency string) ([]*ExchangeRate, error) {
	endpoint := fmt.Sprintf(exchangeRateListEndpointFmt, currency)
	// GET request with no body, so payload is nil
	res, err := c.fetch("GET", endpoint, nil, c.paymentApiKey)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := &exchangeRateListRawResponse{}
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Result, nil
}
