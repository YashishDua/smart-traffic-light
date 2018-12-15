package util

import (
  "encoding/json"
  "net/http"
)

type CustomFunction = func(*http.Request) (interface{}, *HTTPError)

func Response(w http.ResponseWriter, payload interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, err *HTTPError) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.WriteHeader(err.StatusCode)
  body := map[string]string{
    "error": err.Message,
  }
  json.NewEncoder(w).Encode(body)
}

func ResponseWrapper(httpFunction CustomFunction) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    payload, err := httpFunction(r)
    if err != nil {
      Error(w, err)
      return
    }
    Response(w, payload)
  }
}
