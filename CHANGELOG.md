# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Fixed

- Fixed typo in constant name: `createInvoiceEndpoit` → `createInvoiceEndpoint`
- Added HTTP status code validation in `fetch()` method
- Added proper API error response parsing and handling
- Fixed resource leak potential by ensuring response bodies are closed on decode errors
- Improved error messages: "you should pass" → "you must provide one of"
- Fixed `GetBalance()` to return empty struct instead of `nil, nil`
- Fixed test configuration to use environment variables instead of hardcoded placeholders

### Added

- Added `ErrorResponse` type for structured API error handling
- Added validation in `New()` constructor for nil/empty parameters
- Added package-level documentation with usage examples
- Added comprehensive README.md with installation and usage guide
- Added missing test files:
  - `tests/client_test.go` - Client initialization validation tests
  - `tests/webhook_test.go` - Webhook parsing tests
  - `tests/payout_test.go` - Payout validation tests
  - `tests/refund_test.go` - Refund validation tests
  - `tests/other_test.go` - Balance, rates, and services tests
  - `tests/sign_test.go` - Signature verification tests
- Added validation tests for all methods requiring UUID or OrderId
- Added `timeFormat` constant to avoid duplication

### Changed

- Updated `New()` function to return `(*Heleket, error)` instead of `*Heleket`
- Updated examples to handle new error return from `New()`
- Updated test setup to skip tests gracefully when credentials are missing
- Improved error wrapping with context throughout the codebase

### Removed

- Removed `output.md` from version control (already in .gitignore)
