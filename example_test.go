package dydxindexer_test

import (
	"context"

	dydxindexer "github.com/psyb0t/go-dydx-indexer"
)

// This is the README usage example. It talks to the live mainnet indexer, so
// it has no Output comment and only has to compile.
//
//nolint:testableexamples // needs network access to the live API
func Example() {
	client, err := dydxindexer.NewClientWithResponses(
		"https://indexer.dydx.trade/v4",
	)
	if err != nil {
		panic(err)
	}

	res, err := client.GetHeightWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
