package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wkharold/uber-apps/tasks/cmd/app/data"

	"golang.org/x/net/context"
)

type tasktest struct {
	description string
	hfn         ContextHandlerFunc
	req         string
	ctx         context.Context
	rc          int
	body        string
}

var tt = []tasktest{
	{"empty task list", tasklist, "/tasks", notasks(), 200, data.Emptylist},
	{"single task", tasklist, "/tasks", onetask(), 200, data.Singletask},
	{"multiple tasks", tasklist, "/tasks", multipletasks(), 200, data.Multipletasks},
}

func TestTasks(t *testing.T) {
	for _, tst := range tt {
		req, err := http.NewRequest("GET", tst.req, nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		tst.hfn(tst.ctx, w, req)

		if w.Code != tst.rc {
			t.Errorf("%s: Response Code mismatch: expected %d, got %d", tst.description, tst.rc, w.Code)
			continue
		}

		if w.Body.String() != tst.body {
			t.Errorf("%s: Body mismatch: expected %s, got %s", tst.description, tst.body, w.Body.String())
			continue
		}
	}
}

func notasks() context.Context {
	return context.Background()
}

func onetask() context.Context {
	return context.Background()
}

func multipletasks() context.Context {
	return context.Background()
}
