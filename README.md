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
      goat.WriteError(w, 500, "Not implemented")
}

func helloHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
      goat.WriteJSON(w, map[string]string{
          "hello": p["name"],
      })
}

func main() {
      r := goat.New()

      r.Get("/", "", r.IndexHandler)
      r.Get("/hello/:name", "hello_url", helloHandler)

      sr := r.Subrouter("/user")
      sr.Get("/login", "user_login_url", notImplementedHandler)
      sr.Get("/logout", "user_logout_url", notImplementedHandler)

      http.ListenAndServe(":8080", r)
}
```

## Features
### Groups
You can group the routes by a prefix. This can have a serious impact on the
readability of your code.

### Indices
Every route can have a description (like `user_login_url`). These can be used
to automagically generate an API index (like [this](https://api.github.com)).
If you want to hide specific methods, just provide an empty string.

**Note:** Indices are only supported for `GET` requests. Open an issue, if you
want them on other methods, too

### Helpers
You can quickly pretty print JSON to a `http.ResponseWriter` using
`goat.WriteJSON` or `goat.WriteError`.

## Roadmap
* [x] Subrouters or Grouping
* [ ] Middleware
* [ ] Continous integration

## Credits
Goat uses the blazing fast
[httprouter](https://github.com/julienschmidt/httprouter) from Julien Schmidt,
you should really try it out.


