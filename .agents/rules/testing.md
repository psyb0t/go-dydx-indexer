# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `dydxindexer_test`, and use only the exported API.
- **TST2** REST tests run against `httptest.NewServer` and assert the HTTP method, the path, the query parameters, that no `Authorization` header is sent, and the decoded response fields. No test calls the live indexer.
- **TST3** Tests call `t.Parallel()` and use `require` for preconditions and `assert` for values.
- **TST4** `example_test.go` holds the README usage example. When the README example changes, the example changes in the same commit, so `make test` keeps proving it compiles.
- **TST5** Test data uses obvious placeholders such as `TICKER-X`, never real addresses or account data.
