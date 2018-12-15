package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "github.com/go-chi/chi"
  "smart-signal/util"
  "smart-signal/endpoints"
)

func main() {
  router := chi.NewRouter()
  port := os.Getenv("PORT")
  if port == "" {
    fmt.Errorf("$PORT not set")
    return
  }
  log.Println(port)
  router.Get("/", util.ResponseWrapper(Hello))
  router.Post("/traffic/data", util.ResponseWrapper(endpoints.GetTrafficData))
  log.Fatal(http.ListenAndServe(":" + port, router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
