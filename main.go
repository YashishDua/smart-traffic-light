package main

import (
  "log"
  "net/http"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "smart-signal/util"
  "smart-signal/endpoints"
)

func main() {
  router := chi.NewRouter()
  router.Use(middleware.Logger)
  router.Get("/", util.ResponseWrapper(Hello))
  router.Get("/traffic/data", util.ResponseWrapper(endpoints.GetTrafficData))
  log.Fatal(http.ListenAndServe(":8000", router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
