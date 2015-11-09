package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/context"
)

type tasktest struct {
	hfn  ContextHandlerFunc
	req  string
	ctx  context.Context
	rc   int
	body string
}

var tt = []tasktest{
	{tasklist, "/task", notasks(), 200, ""},
}

func TestTasks(t *testing.T) {
	for _, tst := range tt {
		req, err := http.NewRequest("GET", tst.req, nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		tst.hfn(tst.ctx, w, req)

		if w.Code != tst.rc {
			t.Fatalf("Response Code mismatch: expected %d, got %d", tst.rc, w.Code)
		}

		if w.Body.String() != tst.body {
			t.Fatalf("Body mismatch: expected %s, got %s", tst.body, w.Body.String())
		}
	}
}

func notasks() context.Context {
	return context.Background()
}
