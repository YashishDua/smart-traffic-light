package main

import (
  "log"
  "net/http"
  "fmt"
  "os"
  "github.com/go-chi/chi"
  "github.com/go-chi/cors"
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
  cors := cors.New(cors.Options{
    // AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  })
  router.Use(cors.Handler)
  router.Get("/", util.ResponseWrapper(Hello))
  router.Post("/traffic/data", util.ResponseWrapper(endpoints.GetTrafficData))
  log.Fatal(http.ListenAndServe(":" + port, router))
}

//HealhCheck API
func Hello(r *http.Request) (interface{}, *util.HTTPError) {
  return "Under Construction", nil
}
