package main

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

type message struct {
	Data string `json:"data"`
}

func First(client http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "http://localhost:8080/second", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var msg message
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msg.Data += " world"
		err = json.NewEncoder(w).Encode(msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func Second() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span := trace.SpanFromContext(r.Context())
		span.AddEvent("some random event")
		defer span.End()

		err := json.NewEncoder(w).Encode(&message{Data: "Hello"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
