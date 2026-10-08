package api

import (
	"net/http"
)

const helpText = `listkit — agentic-first list and array manipulation service

The agent IS the interface. No UI, no SDK. Send lists as plain text in the
request body, one item per line. Get results as plain text, one item per line.
Add ?format=json or Accept: application/json for JSON output.

INPUT FORMAT
  Single list: items in the body, one per line.
  Multiple lists: lists separated by a line containing only ---.

AUTH
  This service is stateless and requires no authentication by default.
  Use -no-auth flag or LISTKIT_NO_AUTH=1 to explicitly disable auth.

ENDPOINTS

  POST /sort?order=asc|desc&mode=auto|numeric|text
    Sort items. order defaults to asc, mode defaults to auto.
    Example: curl -d 'banana\napple\ncherry' localhost:8470/sort
    Output: apple\nbanana\ncherry

  POST /unique          Remove duplicates, preserve order
  POST /dedup           Alias for /unique
  POST /reverse         Reverse the list
  POST /shuffle         Random shuffle (Fisher-Yates)
  POST /head?n=5        First N items (negative = all but last |N|)
  POST /tail?n=5        Last N items (negative = all but first |N|)
  POST /take?n=5        Alias for /head
  POST /drop?n=2        Drop first N items
  POST /chunk?size=3    Split into chunks of N (output: chunks separated by ---)
  POST /flatten?sep=,   Flatten items by splitting on separator
  POST /compact         Remove empty/whitespace-only items
  POST /rotate?n=2      Rotate by N (positive=left, negative=right)
  POST /sample?n=3      N random items without replacement
  POST /slice?start=1&end=4  Slice from start to end (exclusive, 0-indexed)
  POST /pad?item=x&n=10     Pad list to length N with item
  GET  /fill?item=x&n=5     Create list of N copies of item
  GET  /repeat?item=x&n=5   Alias for /fill
  GET  /range?start=1&end=10&step=1  Generate number range
  POST /transpose       Transpose tab-separated rows to columns
  POST /filter?pattern=foo&mode=contains|prefix|suffix|exact|regex
    Filter items by pattern. Default mode: contains.
  POST /grep            Alias for /filter
  POST /replace?from=old&to=new  Replace text in each item
  POST /join?sep=,      Join items into a single string
  POST /split?sep=,     Split a string into a list
  POST /contains?item=x Check if list contains item (returns true/false)
  POST /index?item=x    Find 0-based index of item (returns -1 if not found)
  POST /count           Count total items

SET OPERATIONS (multiple lists separated by ---)
  POST /intersect       Items in ALL lists
  POST /union           All unique items across lists
  POST /difference      Items in first list not in subsequent lists
  POST /symdiff         Items in exactly one list
  POST /zip             Interleave items from multiple lists

NUMERIC OPERATIONS
  POST /sum             Sum all numeric values
  POST /avg             Average of numeric values
  POST /median          Median of numeric values
  POST /mode             Most frequent value(s)
  POST /min?mode=auto|numeric|text   Minimum value
  POST /max?mode=auto|numeric|text   Maximum value
  POST /frequency       Frequency count (item count per line)
  POST /stats           Summary statistics (JSON: count, unique, min, max, sum, avg, median, mode, range)

REGEX OPERATIONS
  POST /find?pattern=\w+           Find all regex matches in each item
  POST /regex-replace?pattern=old&replacement=new  Replace regex matches

MCP
  POST /mcp             Model Context Protocol (JSON-RPC 2.0)
    Methods: initialize, tools/list, tools/call
    Tools: sort, unique, reverse, shuffle, head, tail, chunk, flatten,
    compact, rotate, sample, slice, pad, fill, range, transpose, filter,
    replace, join, split, contains, index, count, intersect, union,
    difference, symdiff, zyp, sum, avg, median, mode, min, max, frequency,
    stats, find, regex_replace

ERRORS
  Errors are plain text: error: message | hint: what to do next
  Example: error: missing pattern parameter | hint: add ?pattern=foo

CONFIG
  -addr string     Listen address (env: LISTKIT_ADDR, default :8470)
  -no-auth         Disable auth (env: LISTKIT_NO_AUTH)

BUILD
  CGO_ENABLED=0 go build -trimpath ./cmd/listkit
  make build / make test / make vet
`

func (h *Handler) help(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(helpText))
}
