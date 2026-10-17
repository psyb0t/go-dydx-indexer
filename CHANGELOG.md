# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `dydxindexer`, a REST client generated with oapi-codegen from the dYdX v4 Indexer OpenAPI document (Indexer API, spec version v1.0.0, OpenAPI 3.0.0, from `indexer/services/comlink/public/swagger.json` in dydxprotocol/v4-chain). It covers all 49 paths: addresses and subaccounts, affiliates, asset and perpetual positions, candles, compliance screening, fills, funding, historical PnL, orderbooks, orders, perpetual markets, rewards, sparklines, trades, transfers, Turnkey sign-in and vaults.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, renames the two operation IDs the spec uses twice (`GetAffiliateAddress`, `GetComplianceScreen`), and regenerates the client. The spec is not committed.
- Added the `go-dydx-indexer` agent skill. Install it with `claude plugin install go-dydx-indexer@psyb0t` (Claude Code), `codex plugin add go-dydx-indexer@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-dydx-indexer` (ClawHub).
- Licensed under AGPL-3.0, because the generated code derives from dYdX's AGPL-3.0 specification. dYdX's license, with its additional terms, is in `LICENSE.dydx`, and `NOTICE` describes what the pipeline changes.
