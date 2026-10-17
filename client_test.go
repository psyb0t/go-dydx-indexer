package dydxindexer_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	dydxindexer "github.com/psyb0t/go-dydx-indexer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *dydxindexer.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	client, err := dydxindexer.NewClientWithResponses(srv.URL)
	require.NoError(t, err)

	return client
}

func TestGetHeight(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/height", r.URL.Path)
		assert.Empty(t, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"height":"12345","time":"2026-01-02T03:04:05.000Z"}`,
		))
	})

	res, err := client.GetHeightWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, "12345", res.JSON200.Height)
	assert.Equal(t, "2026-01-02T03:04:05.000Z", res.JSON200.Time)
}

func TestGetCandles(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/candles/perpetualMarkets/TICKER-X", r.URL.Path)
		assert.Equal(t, "1HOUR", r.URL.Query().Get("resolution"))
		assert.Equal(t, "2", r.URL.Query().Get("limit"))
		assert.Empty(t, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"candles":[{"id":"c1","ticker":"TICKER-X",` +
				`"resolution":"1HOUR","open":"10.5","close":"11.25",` +
				`"high":"12","low":"10"}]}`,
		))
	})

	limit := 2.0

	res, err := client.GetCandlesWithResponse(
		context.Background(),
		"TICKER-X",
		&dydxindexer.GetCandlesParams{
			Resolution: dydxindexer.CandleResolutionN1HOUR,
			Limit:      &limit,
		},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Candles, 1)

	candle := res.JSON200.Candles[0]
	assert.Equal(t, "c1", candle.Id)
	assert.Equal(t, "10.5", candle.Open)
	assert.Equal(t, "11.25", candle.Close)
}
