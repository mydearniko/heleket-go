# Heleket Go SDK

Official Go client library for the [Heleket](https://heleket.com) cryptocurrency payment gateway API.

## Features

- **Payment Invoices**: Create and manage cryptocurrency payment invoices
- **Payouts**: Send cryptocurrency to any address
- **Static Wallets**: Generate persistent wallet addresses for users
- **Refunds**: Process refunds for payments and blocked wallets
- **Webhooks**: Receive and verify payment notifications
- **Balance & Rates**: Query account balances and exchange rates

## Installation

```bash
go get github.com/idanyas/heleket-go
```

## Quick Start

```go
package main

import (
    "log"
    "net/http"
    
    "github.com/idanyas/heleket-go"
)

func main() {
    client, err := heleket.New(
        &http.Client{},
        "your-merchant-id",
        "your-payment-api-key",
        "your-payout-api-key",
    )
    if err != nil {
        log.Fatal(err)
    }
    
    // Create a payment invoice
    invoice, err := client.CreateInvoice(&heleket.InvoiceRequest{
        Amount:   "10.50",
        Currency: "USD",
        OrderId:  "order-123",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Payment URL: %s", invoice.Url)
}
```

## Configuration

Set your API credentials as environment variables:

```bash
export HELEKET_MERCHANT_ID="your-merchant-id"
export HELEKET_PAYMENT_API_KEY="your-payment-api-key"
export HELEKET_PAYOUT_API_KEY="your-payout-api-key"
```

## Usage Examples

### Create Payment Invoice

```go
invoice, err := client.CreateInvoice(&heleket.InvoiceRequest{
    Amount:   "10.50",
    Currency: "USD",
    OrderId:  "order-123",
    InvoiceRequestOptions: &heleket.InvoiceRequestOptions{
        ToCurrency:  "USDT",
        Network:     "tron",
        UrlCallback: "https://your-site.com/webhook",
        Lifetime:    3600, // 1 hour
    },
})
```

### Get Payment Info

```go
payment, err := client.GetPaymentInfo(&heleket.PaymentInfoRequest{
    PaymentUUID: "payment-uuid",
    // OR
    OrderId: "order-123",
})
```

### Create Payout

```go
payout, err := client.CreatePayout(&heleket.PayoutRequest{
    Amount:     "5.00",
    Currency:   "USDT",
    Network:    "tron",
    OrderId:    "payout-456",
    Address:    "TXYZabc123...",
    IsSubtract: true,
})
```

### Create Static Wallet

```go
wallet, err := client.CreateStaticWallet(&heleket.StaticWalletRequest{
    Currency: "USDT",
    Network:  "tron",
    OrderId:  "user-wallet-789",
    StaticWalletRequestOptions: &heleket.StaticWalletRequestOptions{
        UrlCallback: "https://your-site.com/webhook",
    },
})
```

### Verify Webhook Signature

```go
webhook, err := client.ParseWebhook(requestBody, true) // true = verify signature
if err != nil {
    log.Printf("Invalid webhook: %v", err)
    return
}

switch webhook.Type {
case "payment":
    log.Printf("Payment received: %s", webhook.UUID)
case "payout":
    log.Printf("Payout processed: %s", webhook.UUID)
case "wallet":
    log.Printf("Wallet payment: %s", webhook.UUID)
}
```

### Get Account Balance

```go
balance, err := client.GetBalance()
if err != nil {
    log.Fatal(err)
}

for _, wallet := range balance.Merchant {
    log.Printf("%s: %s", wallet.CurrencyCode, wallet.Balance)
}
```

## API Documentation

For detailed API documentation, visit [Heleket API Docs](https://heleket.com/docs).

## Examples

See the `examples/` directory for complete working examples:

- `examples/payments.go` - Payment invoice operations
- `examples/payouts.go` - Payout operations
- `examples/static_wallets.go` - Static wallet management
- `examples/refunds.go` - Refund processing
- `examples/webhooks.go` - Webhook handling
- `examples/other_features.go` - Balance, rates, discounts

Run examples:

```bash
go run examples/*.go
```

## Testing

Run tests with your API credentials:

```bash
export HELEKET_MERCHANT_ID="your-merchant-id"
export HELEKET_PAYMENT_API_KEY="your-payment-api-key"
export HELEKET_PAYOUT_API_KEY="your-payout-api-key"

go test ./tests/...
```

## Error Handling

The SDK returns detailed error information:

```go
invoice, err := client.CreateInvoice(req)
if err != nil {
    // Check if it's an API error
    if apiErr, ok := err.(*heleket.ErrorResponse); ok {
        log.Printf("API Error (state=%d): %s", apiErr.State, apiErr.Message)
    } else {
        log.Printf("Request failed: %v", err)
    }
    return
}
```

## Security

- Always use HTTPS in production
- Store API keys securely (environment variables, secrets manager)
- Verify webhook signatures to prevent tampering
- Never commit API keys to version control

## License

MIT License - see LICENSE file for details

## Support

- Documentation: https://heleket.com/docs
- Issues: https://github.com/idanyas/heleket-go/issues
- Email: support@heleket.com

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
