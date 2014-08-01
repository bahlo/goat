# Goat

Goat is a Go REST server. You can pronounce it like the _goat_, or
_go-at_.

## Usage
```go
package main

import (
        "net/http"

        "github.com/bahlo/goat"
)

func main() {
        r := goat.New()

        r.Get("/user", "user_url", func(w http.ResponseWriter, r *http.Request) {
                goat.WriteError(w, "Not implemented")
        }

        http.ListenAndServe(":8080", r)
}
```

## Features
* Speed (more info on the way)
* Helper functions for working with JSON (`WriteError`, `WriteJSON`)
* An index generator to get an index like [this](https://api.github.com/)
* Optimized for machines: No templates, static files, etc. (if you want this,
  you'd be better off with one of
  [these](https://github.com/avelino/awesome-go#web-frameworks)

## Credits
Goat uses the blazing fast
[httprouter](https://github.com/julienschmidt/httprouter) from Julien Schmidt,
you should really try it out.


