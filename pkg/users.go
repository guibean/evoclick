package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	Path         string     `json:"path"`
	Method       string     `json:"method"`
	QueryStrings url.Values `json:"query_strings"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "  ")

	debugRequest := &Request{
		Path:         r.URL.Path,
		QueryStrings: r.URL.Query(),
		Method:       r.Method,
	}

	err := jsonEncoder.Encode(debugRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
