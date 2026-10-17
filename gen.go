// Package dydxindexer is a Go client for the dYdX v4 Indexer HTTP API,
// generated with oapi-codegen from the OpenAPI document in the dYdX v4-chain
// repository.
package dydxindexer

//go:generate bash scripts/fetch-spec.sh
//go:generate bash scripts/fix-operation-ids.sh
//go:generate go tool oapi-codegen --config=config.yaml api/openapi.fixed.json
