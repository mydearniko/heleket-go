# Contributing to Heleket Go SDK

Thank you for your interest in contributing to the Heleket Go SDK!

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/idanyas/heleket-go.git
cd heleket-go
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your test environment variables:
```bash
export HELEKET_MERCHANT_ID="your-test-merchant-id"
export HELEKET_PAYMENT_API_KEY="your-test-payment-key"
export HELEKET_PAYOUT_API_KEY="your-test-payout-key"
```

## Running Tests

Run all tests:
```bash
go test ./tests/...
```

Run tests with verbose output:
```bash
go test -v ./tests/...
```

Run specific test:
```bash
go test -v ./tests/ -run TestCreateInvoice
```

## Code Quality

Before submitting a PR, ensure your code passes:

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Run tests
go test ./tests/...
```

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to your branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Coding Standards

- Follow standard Go conventions and idioms
- Add godoc comments for all exported functions and types
- Keep functions focused and single-purpose
- Handle errors explicitly
- Use meaningful variable names
- Add tests for new features and bug fixes

## Reporting Issues

When reporting issues, please include:
- Go version (`go version`)
- SDK version
- Minimal code to reproduce the issue
- Expected vs actual behavior
- Any error messages or logs

Thank you for contributing!
