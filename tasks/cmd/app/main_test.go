package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/wkharold/uber-apps/tasks/cmd/app/data"

	"golang.org/x/net/context"
)

const (
	GET  = "GET"
	POST = "POST"
)

type tasktest struct {
	description string
	hfn         ContextHandlerFunc
	req         string
	method      string
	payload     string
	ctx         context.Context
	rc          int
	body        string
}

var tt = []tasktest{
	{"empty task list", tasklist, "/tasks", GET, "", notasks(), 200, data.Emptylist},
	{"single task", tasklist, "/tasks", GET, "", onetask(), 200, data.Singletask},
	{"multiple tasks", tasklist, "/tasks", GET, "", multipletasks(), 200, data.Multipletasks},
	{"add task to empty list", taskadd, "/tasks", POST, "text=another task", notasks(), 204, ""},
	{"add task to existing tasks", taskadd, "/tasks", POST, "text=another task", multipletasks(), 204, ""},
	{"bad add request", taskadd, "/tasks", POST, "task=another task", multipletasks(), 400, ""},
}

func TestTasks(t *testing.T) {
	for _, tst := range tt {
		req, err := http.NewRequest(tst.method, tst.req, strings.NewReader(tst.payload))
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		tst.hfn(tst.ctx, w, req)

		if w.Code != tst.rc {
			t.Errorf("%s: Response Code mismatch: expected %d, got %d", tst.description, tst.rc, w.Code)
			continue
		}

		if len(tst.body) == 0 {
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
	ctx := context.WithValue(context.Background(), "tasks", list.New())
	return ctx
}

func onetask() context.Context {
	l := list.New()
	l.PushBack("task one")

	ctx := context.WithValue(context.Background(), "tasks", l)
	return ctx
}

func multipletasks() context.Context {
	l := list.New()
	l.PushBack("task one")
	l.PushBack("task two")
	l.PushBack("task three")

	ctx := context.WithValue(context.Background(), "tasks", l)
	return ctx
}
