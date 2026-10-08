package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// wantsJSON checks if the client wants JSON output.
func wantsJSON(r *http.Request) bool {
	if r.URL.Query().Get("format") == "json" {
		return true
	}
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "application/json")
}

// writeList writes a list of items as plain text (one per line) or JSON.
func writeList(w http.ResponseWriter, r *http.Request, items []string) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for _, item := range items {
		w.Write([]byte(item))
		w.Write([]byte("\n"))
	}
}

// writeMultiList writes multiple lists separated by --- in plain text, or JSON.
func writeMultiList(w http.ResponseWriter, r *http.Request, lists [][]string) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lists)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for i, list := range lists {
		if i > 0 {
			w.Write([]byte("---\n"))
		}
		for _, item := range list {
			w.Write([]byte(item))
			w.Write([]byte("\n"))
		}
	}
}

// writeString writes a single string value.
func writeString(w http.ResponseWriter, r *http.Request, val string) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": val})
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(val))
	w.Write([]byte("\n"))
}

// writeBool writes a boolean value.
func writeBool(w http.ResponseWriter, r *http.Request, val bool) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"result": val})
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if val {
		w.Write([]byte("true\n"))
	} else {
		w.Write([]byte("false\n"))
	}
}

// writeInt writes an integer value.
func writeInt(w http.ResponseWriter, r *http.Request, val int) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"result": val})
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(itoa(val)))
	w.Write([]byte("\n"))
}

// writeError writes an error response with a hint.
func writeError(w http.ResponseWriter, code int, msg, hint string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte("error: "))
	w.Write([]byte(msg))
	if hint != "" {
		w.Write([]byte(" | hint: "))
		w.Write([]byte(hint))
	}
	w.Write([]byte("\n"))
}

// writeJSON writes an arbitrary JSON response.
func writeJSON(w http.ResponseWriter, r *http.Request, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}

// writeFreq writes a frequency map.
func writeFreq(w http.ResponseWriter, r *http.Request, freq map[string]int) {
	if wantsJSON(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(freq)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	for k, v := range freq {
		w.Write([]byte(k))
		w.Write([]byte(" "))
		w.Write([]byte(itoa(v)))
		w.Write([]byte("\n"))
	}
}

// itoa is a simple int to string converter.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
