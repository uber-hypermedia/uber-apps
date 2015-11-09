// Package main provides a simple UBER hypermedia drive todo list server
package main

import (
	"net/http"

	"golang.org/x/net/context"
)

type ContextHandler interface {
	ServeHTTPWithContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPWithContext(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	h.ServeHTTPWithContext(ctx, w, req)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func (ca ContextAdapter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPWithContext(ca.ctx, w, req)
}

func main() {
	panic("not implemented")
}

func tasklist(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
