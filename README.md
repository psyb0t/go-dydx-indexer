# go-dydx-indexer

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-dydx-indexer.svg)](https://pkg.go.dev/github.com/psyb0t/go-dydx-indexer)
[![CI](https://github.com/psyb0t/go-dydx-indexer/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-dydx-indexer/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-dydx-indexer/badges/version.svg)](https://github.com/psyb0t/go-dydx-indexer/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-dydx-indexer/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-dydx-indexer/badges/importers.svg)](https://github.com/psyb0t/go-dydx-indexer/blob/badges/importers.md)

Unofficial Go client for the [dYdX v4 Indexer API](https://docs.dydx.xyz), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document in the [dYdX v4-chain repository](https://github.com/dydxprotocol/v4-chain). I am not affiliated with dYdX.

The package is `dydxindexer`, in the repository root. Everything in `client.gen.go` is generated. Don't edit it by hand, regenerate it.

## Install

```sh
go get github.com/psyb0t/go-dydx-indexer
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"

	dydxindexer "github.com/psyb0t/go-dydx-indexer"
)

func main() {
	client, err := dydxindexer.NewClientWithResponses("https://indexer.dydx.trade/v4")
	if err != nil {
		panic(err)
	}

	res, err := client.GetHeightWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
```

The same code is in `example_test.go`, so `make test` checks that it compiles.

The full API, with the spec's descriptions as doc comments, is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-dydx-indexer) and in `go doc github.com/psyb0t/go-dydx-indexer`.

Method names come from the spec's operation IDs: `GetCandles`, `GetFills`, `ListPerpetualMarkets`, `GetPerpetualMarket` (the orderbook) and so on. The spec uses two IDs twice, so `make generate` renames the second of each: `GetAffiliateAddress` is `GET /affiliates/address` and `GetComplianceScreen` is `GET /compliance/screen/{address}`. `GetAddress` and `Screen` keep their original paths. `Get` is the sparklines endpoint.

Prices, sizes and other decimal values come back as strings, as the indexer sends them.

## Auth

The Indexer API is public and needs no key. The base URL includes the `/v4` prefix, because the paths in the spec start after it:

| Network | Base URL |
|---|---|
| Mainnet | `https://indexer.dydx.trade/v4` |
| Testnet | `https://indexer.v4testnet.dydx.exchange/v4` |

The dYdX docs list the testnet HTTP endpoint as `https://indexer.v4testnet.dydx.exchange`; add `/v4` to it as shown above.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `indexer/services/comlink/public/swagger.json` from the `main` branch of dydxprotocol/v4-chain to `api/openapi.json`.
- `scripts/fix-operation-ids.sh` writes `api/openapi.fixed.json`. The spec gives `GET /addresses/{address}` and `GET /affiliates/address` the same operation ID `GetAddress`, and `GET /screen` and `GET /compliance/screen/{address}` the same ID `Screen`, which oapi-codegen turns into duplicate Go methods. The script renames the second of each pair with jq and fails if a renamed path is gone or any ID is still duplicated.
- oapi-codegen runs with `config.yaml`. The config maps `number` to `float64`, because the spec gives numbers no format and oapi-codegen would otherwise use `float32`.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-dydx-indexer) teaches a coding agent how to use go-dydx-indexer when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-dydx-indexer@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-dydx-indexer@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-dydx-indexer
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[AGPL-3.0](LICENSE). The generated code derives from the Indexer API specification dYdX Trading Inc. publishes in the dYdX v4-chain repository under the AGPL-3.0 with additional terms of its own, which also apply to the generated code; see [LICENSE.dydx](LICENSE.dydx) and [NOTICE](NOTICE). The specification itself is not redistributed here.
