# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.2, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.
- Raised the minimum Go version in `go.mod` from 1.26.9 to 1.26.10, which fixes the `crypto/x509` vulnerability GO-2026-6809 found by govulncheck.
- The issue-pull workflow runs at hours derived from the repository name instead of the same four hours as every other repository.

## v0.1.1, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.

## v0.1.0, 2026-10-17

- Added `dydxindexer`, a REST client generated with oapi-codegen from the dYdX v4 Indexer OpenAPI document (Indexer API, spec version v1.0.0, OpenAPI 3.0.0, from `indexer/services/comlink/public/swagger.json` in dydxprotocol/v4-chain). It covers all 49 paths: addresses and subaccounts, affiliates, asset and perpetual positions, candles, compliance screening, fills, funding, historical PnL, orderbooks, orders, perpetual markets, rewards, sparklines, trades, transfers, Turnkey sign-in and vaults.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, renames the two operation IDs the spec uses twice (`GetAffiliateAddress`, `GetComplianceScreen`), and regenerates the client. The spec is not committed.
- Added the `go-dydx-indexer` agent skill. Install it with `claude plugin install go-dydx-indexer@psyb0t` (Claude Code), `codex plugin add go-dydx-indexer@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-dydx-indexer` (ClawHub).
- Licensed under AGPL-3.0, because the generated code derives from dYdX's AGPL-3.0 specification. dYdX's license, with its additional terms, is in `LICENSE.dydx`, and `NOTICE` describes what the pipeline changes.
