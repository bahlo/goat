# Goat

Goat is a Go REST server. You can pronounce it like the _goat_, or
_go-at_. Depends on how you like goats.

## Usage
```go
package main

import (
    "net/http"

    "github.com/bahlo/goat"
)

func notImplementedHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
      goat.WriteError(w, "Not implemented")
}

func helloHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
      goat.WriteJSON(w, map[string]string{
          "hello": p["name"],
      })
}

func main() {
      r := goat.New()

      r.Get("/user", "user_url", notImplementedHandler)
      r.Get("/hello/:name", "hello_url", helloHandler)

      http.ListenAndServe(":8080", r)
}
```

## Roadmap
* [ ] Subrouters or Grouping
* [ ] Middleware
* [ ] Continous integration

## Credits
Goat uses the blazing fast
[httprouter](https://github.com/julienschmidt/httprouter) from Julien Schmidt,
you should really try it out.


