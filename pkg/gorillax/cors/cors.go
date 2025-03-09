// Copyright (c) 2013, SoundCloud Ltd.
// Use of this source code is governed by a BSD-style
// license that can be found in the README file.
// Source code and contact info at http://github.com/streadway/handy

/*
Package cors contains filters to handle CORS related requests defined from
http://www.w3.org/TR/cors/
*/

// copy from github.com/streadway/handy/cors

package cors

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// options parameterizes CORS behavior.
type options struct {
	// AllowOrigin transforms a request into the Access-Control-Allow-Origin
	// header, default is full access "*".
	AllowOrigin func(*http.Request) string
	MaxAge      time.Duration
}

type Option func(*options)

func AllowOrigin(origin func(*http.Request) string) Option {
	return func(cfg *options) {
		cfg.AllowOrigin = origin
	}
}

// Middleware returns a middleware that applies options to the request.
func Middleware(opts ...Option) mux.MiddlewareFunc {
	opt := &options{
		AllowOrigin: func(request *http.Request) string {
			// DefaultAllowOrigin sets Access-Control-Allow-Origin when not configured.
			return "*"
		},
		MaxAge: 10 * time.Minute,
	}

	age := strconv.Itoa(int(opt.MaxAge / time.Second))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Encoding, Authorization, Content-Type, Origin")
			w.Header().Set("Access-Control-Allow-Origin", opt.AllowOrigin(r))

			switch r.Method {
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			case "OPTIONS":
				if r.Header.Get("Access-Control-Request-Method") == "GET" {
					w.Header().Set("Access-Control-Max-Age", age)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
			case "HEAD", "GET":
				next.ServeHTTP(w, r)
			}
		})
	}
}

// Get implements a simple read-only access control policy handling preflight
// and normal requests with a cache age of 10 minutes for preflight requests.
// Methods other than HEAD, OPTIONS, GET will return 405.
//
// The origin parameter should be the case-insentive fully qualified origin
// domain to match or '*' to match any domain.
func Get(origin string, next http.Handler) http.Handler {
	return Middleware(AllowOrigin(func(request *http.Request) string { return origin }))(next)
}
