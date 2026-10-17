---
name: go-dydx-indexer
description: Unofficial Go client for the dYdX v4 Indexer HTTP API, generated with oapi-codegen. `dydxindexer.NewClientWithResponses(baseURL)` from `github.com/psyb0t/go-dydx-indexer` gives a typed method for every endpoint in the indexer's OpenAPI spec (block height and time, perpetual markets, orderbooks, candles, trades, historical funding, sparklines, subaccounts, positions, orders, fills, transfers, PnL, rewards, affiliates, vaults, compliance screening), with decoded `JSON200` response fields. The API is public and needs no key. Use when Go code needs dYdX v4 market or account data from the indexer, or already imports go-dydx-indexer.
homepage: https://github.com/psyb0t/go-dydx-indexer
user-invocable: true
metadata:
  openclaw:
    emoji: "📈"
    requires:
      bins:
        - go
permissions:
  filesystem: "Guides edits to the Go source of the current project. The library reads and writes no files."
  shell: "The Go toolchain (`go get`, `go build`, `go test`) to add and compile the dependency."
  network: "`go get` fetches the module from your Go module proxy. At runtime the client sends HTTPS requests only to the base URL you pass, normally https://indexer.dydx.trade/v4 (mainnet) or https://indexer.v4testnet.dydx.exchange/v4 (testnet)."
---

# go-dydx-indexer

Typed Go client for the dYdX v4 Indexer HTTP API. Package `dydxindexer`, generated from the OpenAPI document in the dydxprotocol/v4-chain repository.

For install details, the operation list and response types, see [references/setup.md](references/setup.md).

## Security & safety

- The indexer is public. No API key is needed, so none should be added to requests.
- Requests go only to the base URL passed to `NewClientWithResponses`. Use the mainnet or testnet indexer URL from the table in the reference, as the user asks.
- The indexer reads data; it cannot place orders or move funds. Its few POST endpoints (push token registration, Turnkey sign-in, affiliate referral codes) act on the caller's identity, so call them only when the user asked for that action.
- Wallet addresses are public on chain, but treat the ones the user gives as their data: send them only to the indexer.

## When to use

- Go code needs dYdX v4 market data: markets, orderbooks, candles, trades, funding, sparklines.
- Go code needs account data for an address: subaccounts, positions, orders, fills, transfers, PnL, rewards.
- The project already imports `github.com/psyb0t/go-dydx-indexer`.

## When NOT to use

- Placing or cancelling orders, or any other transaction. That goes to a dYdX validator node, not the indexer.
- Live streaming. The indexer WebSocket feed is not in this client.

## Usage

```go
import (
	"context"
	"fmt"

	dydxindexer "github.com/psyb0t/go-dydx-indexer"
)

client, err := dydxindexer.NewClientWithResponses("https://indexer.dydx.trade/v4")
if err != nil {
	return err
}

height, err := client.GetHeightWithResponse(ctx)
if err != nil {
	return err // transport failure
}

if height.JSON200 == nil {
	return fmt.Errorf("get height: status %d: %s", height.StatusCode(), height.Body)
}

fmt.Println(height.JSON200.Height, height.JSON200.Time)
```

Hourly candles for one market:

```go
limit := 24.0

candles, err := client.GetCandlesWithResponse(
	ctx,
	ticker, // a key of ListPerpetualMarkets' JSON200.Markets
	&dydxindexer.GetCandlesParams{
		Resolution: dydxindexer.CandleResolutionN1HOUR,
		Limit:      &limit,
	},
)
if err != nil {
	return err
}

if candles.JSON200 == nil {
	return fmt.Errorf("get candles: status %d: %s", candles.StatusCode(), candles.Body)
}

for _, candle := range candles.JSON200.Candles {
	fmt.Println(candle.StartedAt, candle.Open, candle.Close)
}
```

Prices, sizes and other decimals are strings, as the indexer sends them. Parse them with a decimal library, not `float64`, when precision matters. Numeric query parameters such as `Limit` are `*float64`.

## Errors

- A non-nil `err` is a transport or decoding failure.
- An HTTP error status is not an `err`. Check `StatusCode()`: `JSON200` is set on success, and `Body` always holds the raw response. The spec declares no typed error responses.
