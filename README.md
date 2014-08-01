# Goat

Goat is a Go REST server. You can pronounce it like the _goat_, or
_go-at_. Depends on how you like goats.

## Contents
- [Usage](#usage)
  - [Subrouters](#subrouters)
  - [Indices](#indices)
  - [Middleware](#middleware)
- [Roadmap](#roadmap)
- [Credits](#credits)
- [License](#license)

## Usage
```go
package main

import (
    "net/http"

    "github.com/bahlo/goat"
)

func helloHandler(w http.ResponseWriter, r *http.Request, p goat.Params) {
    goat.WriteJSON(w, map[string]string{
        "hello": p["name"],
    })
}

func main() {
    r := goat.New()

    r.Get("/hello/:name", "hello_url", helloHandler)

    r.Run(":8080")
}
```

### Subrouters
You can create subrouters to simplify your code
```go
func main() {
    r := goat.New()
    r.Get("/hello/:name", "hello_url", helloHandler)

    sr := r.Subrouter("/user")
    {
        sr.Post("/login", "user_login_url", loginHandler)
        sr.Get("/logout", "user_logout_url", logoutHandler)
    }

    r.Run(":8080")
}
```

### Indices
Every route can have a description (like `user_login_url`). These can be used
to automagically generate an API index (like [this](https://api.github.com)).
If you want to hide specific methods, just provide an empty string.

```go
func main() {
    r := goat.New()

    r.Get("/", "", r.IndexHandler)
    r.Get("/hello/:name", "hello_url", helloHandler)

    sr := r.Subrouter("/user")
    {
        sr.Post("/login", "user_login_url", loginHandler)
        sr.Get("/logout", "user_logout_url", logoutHandler)
    }

    r.Run(":8080")
}
```

The above example would return the following response on `/`:
```json
{
  "hello_url": "/hello/:name",
  "user_logout_url": "/user/logout"
}
```

**Note:** Indices are only supported for `GET` requests. Open an issue, if you
want them on other methods, too

### Middleware
You can easily include any middleware you like. Important is, that it's in the
following format:
```go
func(http.Handler) http.Handler
```

Example:
```go
func main() {
    r := goat.New()

    r.Get("/hello/:name", "hello_url", helloHandler)

    r.Use(myMiddleWare, myOtherMiddleWare)

    r.Run(":8080")
}
```

## Roadmap
* [x] Subrouters or Grouping
* [x] Middleware
* [ ] Continous integration

## Philosophy
I wanted to create a small, fast and reliable REST API server, which supports
quick JSON and error output, good rooting and easy-to-use middleware.
I have split the files after responsibility to make it easy for everyone to
dive in (start with `goat.go`.

## Credits
Thanks to Julien Schmidt for the amazing
[httprouter](https://github.com/julienschmidt/httprouter) used in this
project.

## License
This project is licensed unter MIT, for more information look into the LICENSE
file.
