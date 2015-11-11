package main

import (
	"bytes"
	"encoding/json"
	"log"
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

		if equaljson([]byte(w.Body.String()), []byte(tst.body)) == false {
			body := bytes.NewBuffer([]byte{})
			json.Compact(body, []byte(tst.body))
			t.Errorf("%s: Body mismatch:\nexpected %s\ngot      %s", tst.description, string(body.Bytes()), w.Body.String())
			continue
		}
	}
}

func equaljson(p, q []byte) bool {
	cp := bytes.NewBuffer([]byte{})

	if err := json.Compact(cp, p); err != nil {
		log.Printf("unable to compact cp json for equaljson: %+v", err)
		return false
	}

	cq := bytes.NewBuffer([]byte{})

	if err := json.Compact(cq, q); err != nil {
		log.Printf("unable to compact cq json for equaljson: %+v", err)
		return false
	}

	if len(cp.Bytes()) != len(cq.Bytes()) {
		return false
	}

	cpb, cqb := cp.Bytes(), cq.Bytes()

	for i, b := range cpb {
		if b != cqb[i] {
			return false
		}
	}

	return true
}

func notasks() context.Context {
	ctx := context.WithValue(context.Background(), "tasks", []string{})
	return ctx
}

func onetask() context.Context {
	ctx := context.WithValue(context.Background(), "tasks", []string{"task one"})
	return ctx
}

func multipletasks() context.Context {
	ctx := context.WithValue(context.Background(), "tasks", []string{"task one", "task two", "task three"})
	return ctx
}
