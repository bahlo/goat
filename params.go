package goat

import "github.com/julienschmidt/httprouter"

// Params represents the parameters of a request
type Params map[string]string

// paramsFromHTTPRouter converts httprouter.Params to goat.Params
func paramsFromHTTPRouter(hrps httprouter.Params) Params {
	var ps = Params{}

	for _, p := range hrps {
		ps[p.Key] = p.Value
	}

	return ps
}
