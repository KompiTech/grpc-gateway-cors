package cors

import (
	rscors "github.com/rs/cors"
	"net/http"
)

// Wrap wraps an http handler with
func Wrap(h http.Handler, opts ...Option) http.Handler {
	a := &Adapter{}
	for _, opt := range opts {
		opt(a)
	}
	nc := rscors.New(a.Options)
	return nc.Handler(h)
}

// Adapter wraps rs/cors options, but can be changed along
// with options functions to adapt any cors library configurator.
type Adapter struct {
	rscors.Options
}

// Option sets a Adapter configuration option.
type Option func(*Adapter)

// AllowedOrigins sets option for allowed origins.
func AllowedOrigins(os ...string) Option {
	return func(c *Adapter) {
		c.AllowedOrigins = os
	}
}

// AllowedMethods sets options for allowed methods.
func AllowedMethods(ms ...string) Option {
	return func(c *Adapter) {
		c.AllowedMethods = ms
	}
}

// AllowedHeaders sets Adapter options for allowed headers.
func AllowedHeaders(ahs ...string) Option {
	return func(c *Adapter) {
		c.AllowedHeaders = ahs
	}
}

// ExposedHeaders sets Adapter options for exposed headers.
func ExposedHeaders(hs ...string) Option {
	return func(c *Adapter) {
		c.ExposedHeaders = hs
	}
}

// MaxAge sets the max age option.
func MaxAge(i int) Option {
	return func(c *Adapter) {
		c.MaxAge = i
	}
}

// AllowCredentials sets the allow credentials option.
func AllowCredentials(ac bool) Option {
	return func(c *Adapter) {
		c.AllowCredentials = ac
	}
}

// OptionsPassthrough sets the Adapter option method passthrough option.
func OptionsPassthrough(ac bool) Option {
	return func(c *Adapter) {
		c.OptionsPassthrough = ac
	}
}

func Debug(ac bool) Option {
	return func(c *Adapter) {
		c.Debug = ac
	}
}
