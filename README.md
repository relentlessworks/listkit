# listkit

Agentic-first list and array manipulation service. Sort, unique, reverse, shuffle, chunk, flatten, zip, rotate, intersect, union, difference, symmetric difference, filter, sample, head, tail, slice, sum, avg, median, mode, min, max, frequency, range, and more. Plain text API, agent-driven, single Go binary.

## Quick Start

```bash
# Build
make build

# Run (development mode, no auth)
./listkit -no-auth

# Sort a list
curl -d 'banana
apple
cherry' localhost:8470/sort

# Output:
# apple
# banana
# cherry

# Get JSON output
curl -d '3
1
2' -H 'Accept: application/json' localhost:8470/sort?order=desc

# Set intersection (lists separated by ---)
curl -d 'apple
banana
cherry
---
banana
cherry
date' localhost:8470/intersect

# Output:
# banana
# cherry

# Numeric stats
curl -d '1
2
3
4
5' -H 'Accept: application/json' localhost:8470/stats

# Self-documenting help
curl localhost:8470/help
```

## Principles

- **The agent IS the interface** — No UI, no SDK. The API is the product.
- **Plain text by default** — One item per line. JSON on demand via `Accept: application/json` or `?format=json`.
- **Instructive errors** — Every 4xx includes a hint telling the agent what to do next.
- **Self-documenting** — `GET /help` returns a one-page operating manual.
- **Single static binary** — Go, CGO_ENABLED=0, zero external dependencies.
- **Zero config defaults** — Runs out of the box. Config: defaults < env < flags.
- **MCP connector** — Speaks Model Context Protocol at `/mcp`.

## API Reference

### List Operations

| Method | Path | Params | Description |
|--------|------|--------|-------------|
| POST | `/sort` | `order=asc\|desc`, `mode=auto\|numeric\|text` | Sort items |
| POST | `/unique` | — | Remove duplicates |
| POST | `/reverse` | — | Reverse the list |
| POST | `/shuffle` | — | Random shuffle |
| POST | `/head` | `n=10` | First N items |
| POST | `/tail` | `n=10` | Last N items |
| POST | `/drop` | `n=1` | Drop first N items |
| POST | `/chunk` | `size=2` | Split into chunks |
| POST | `/flatten` | `sep=,` | Flatten by separator |
| POST | `/compact` | — | Remove empty items |
| POST | `/rotate` | `n=1` | Rotate by N |
| POST | `/sample` | `n=1` | N random items |
| POST | `/slice` | `start=0`, `end=N` | Slice list |
| POST | `/pad` | `item=x`, `n=10` | Pad to length N |
| GET | `/fill` | `item=x`, `n=5` | Create N copies |
| GET | `/range` | `start=0`, `end=10`, `step=1` | Generate number range |
| POST | `/transpose` | — | Transpose rows to columns |
| POST | `/filter` | `pattern=foo`, `mode=contains` | Filter by pattern |
| POST | `/replace` | `from=old`, `to=new` | Replace text |
| POST | `/join` | `sep=,` | Join into string |
| POST | `/split` | `sep=,` | Split string into list |
| POST | `/contains` | `item=x` | Check if contains |
| POST | `/index` | `item=x` | Find index |
| POST | `/count` | — | Count items |

### Set Operations (multiple lists separated by `---`)

| Method | Path | Description |
|--------|------|-------------|
| POST | `/intersect` | Items in ALL lists |
| POST | `/union` | All unique items |
| POST | `/difference` | First list minus others |
| POST | `/symdiff` | Items in exactly one list |
| POST | `/zip` | Interleave items |

### Numeric Operations

| Method | Path | Description |
|--------|------|-------------|
| POST | `/sum` | Sum numeric values |
| POST | `/avg` | Average |
| POST | `/median` | Median |
| POST | `/mode` | Most frequent value(s) |
| POST | `/min` | Minimum (`mode=auto\|numeric\|text`) |
| POST | `/max` | Maximum (`mode=auto\|numeric\|text`) |
| POST | `/frequency` | Frequency count |
| POST | `/stats` | Summary statistics (JSON) |

### Regex Operations

| Method | Path | Params | Description |
|--------|------|--------|-------------|
| POST | `/find` | `pattern=\w+` | Find all regex matches |
| POST | `/regex-replace` | `pattern=old`, `replacement=new` | Replace regex matches |

### Other

| Method | Path | Description |
|--------|------|-------------|
| GET | `/help` | Self-documenting help |
| POST | `/mcp` | MCP JSON-RPC 2.0 endpoint |

## Configuration

| Flag | Env | Default | Description |
|------|-----|---------|-------------|
| `-addr` | `LISTKIT_ADDR` | `:8470` | Listen address |
| `-no-auth` | `LISTKIT_NO_AUTH` | `false` | Disable auth |

## Build

```bash
make build    # CGO_ENABLED=0 go build -trimpath
make test     # go test -race ./...
make vet      # go vet ./...
```

## License

MIT
