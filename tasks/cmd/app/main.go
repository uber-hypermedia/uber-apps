// Package main provides a simple UBER hypermedia drive todo list server
package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

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

type udata struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Rel        []string `json:"rel,omitempty"`
	Label      string   `json:"label,omitempty"`
	URL        string   `json:"url,omitempty"`
	Template   bool     `json:"template,omitempty"`
	Action     string   `json:"action,omitempty"`
	Transclude bool     `json:"transclude,omitempty"`
	Model      string   `json:"model,omitempty"`
	Sending    string   `json:"sending,omitempty"`
	Accepting  []string `json:"accepting,omitempty"`
	Value      string   `json:"value,omitempty"`
	Data       []udata  `json:"data,omitempty"`
}

type ubody struct {
	Version string  `json:"version"`
	Data    []udata `json:"data,omitempty"`
	Error   []udata `json:"error,omitempty"`
}

type udoc struct {
	Uber ubody `json:"uber"`
}

func main() {
	panic("not implemented")
}

func taskadd(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mkError("ServerError", "reason", "Cannot read HTTP request body"))
		return
	}

	re := regexp.MustCompile("text=(([[:word:]]|[[:space:]])*)")
	sm := re.FindStringSubmatch(string(body))
	if sm == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mkError("ClientError", "reason", "Unrecognized add task body"))
		return
	}

	tasks := ctx.Value("tasks").(*list.List)
	tasks.PushBack(sm[1])

	w.WriteHeader(http.StatusNoContent)
}

func taskcomplete(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mkError("ServerError", "reason", "Cannot read HTTP request body"))
		return
	}

	re := regexp.MustCompile("id=[[:alpha:]]+([[:digit:]]+)")
	sm := re.FindStringSubmatch(string(body))
	if sm == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mkError("ClientError", "reason", "Unrecognized complete text body"))
		return
	}

	completed := false
	taskid, err := strconv.Atoi(sm[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mkError("ServerError", "reason", "Cannot read HTTP request body"))
		return
	}

	tasks := ctx.Value("tasks").(*list.List)

	if tasks.Len() < taskid {
		w.WriteHeader(http.StatusNotFound)
		w.Write(mkError("ClientError", "reason", "No such task"))
		return
	}

	for t, i := tasks.Front(), 1; t != nil; t = t.Next() {
		if i == taskid {
			completed = true
			tasks.Remove(t)
		}
		i++
	}

	if !completed {
		w.WriteHeader(http.StatusNotFound)
		w.Write(mkError("ClientError", "reason", "No such task"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func tasklist(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	tasks := ctx.Value("tasks").(*list.List)

	resp := mkEmptylist()
	if resp == nil {
		panic("can't generate base UBER document")
	}

	for t, i := tasks.Front(), 0; t != nil; t = t.Next() {
		task := udata{ID: fmt.Sprintf("task%d", i+1),
			Rel:  []string{"item"},
			Name: "tasks",
			Data: []udata{
				udata{Rel: []string{"complete"}, URL: "/tasks/complete/", Model: fmt.Sprintf("id=task%d", i+1), Action: "append"},
				udata{Name: "text", Value: t.Value.(string)}}}

		resp.Uber.Data[1].Data = append(resp.Uber.Data[1].Data, task)
		i++
	}

	bs, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mkError("ServerError", "reason", "Cannot read HTTP request body"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func tasksearch(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	tasks := ctx.Value("tasks").(*list.List)

	qt := req.URL.Query().Get("text")
	if len(qt) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mkError("ClientError", "reason", "Missing text parameter"))
		return
	}

	resp := mkEmptylist()
	if resp == nil {
		panic("can't generate base UBER document")
	}

	for t, i := tasks.Front(), 0; t != nil; t = t.Next() {
		if qt == t.Value.(string) {
			task := udata{ID: fmt.Sprintf("task%d", i+1),
				Rel:  []string{"item"},
				Name: "tasks",
				Data: []udata{
					udata{Rel: []string{"complete"}, URL: "/tasks/complete/", Model: fmt.Sprintf("id=task%d", i+1), Action: "append"},
					udata{Name: "text", Value: t.Value.(string)}}}

			resp.Uber.Data[1].Data = append(resp.Uber.Data[1].Data, task)
			i++
		}
	}

	bs, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mkError("ServerError", "reason", "Cannot read HTTP request body"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func mkEmptylist() *udoc {
	links := udata{
		ID: "links",
		Data: []udata{
			udata{ID: "alps",
				Rel:    []string{"profile"},
				URL:    "/tasks-alps.xml",
				Action: "read",
				Data:   []udata{}},
			udata{ID: "list",
				Name:   "links",
				Rel:    []string{"collection"},
				URL:    "/tasks/",
				Action: "read",
				Data:   []udata{}},
			udata{ID: "search",
				Name:   "links",
				Rel:    []string{"search"},
				URL:    "/tasks/search",
				Action: "read",
				Model:  "?text={text}",
				Data:   []udata{}},
			udata{ID: "add",
				Name:   "links",
				Rel:    []string{"add"},
				URL:    "/tasks/",
				Action: "append",
				Model:  "text={text}",
				Data:   []udata{}}}}

	return &udoc{ubody{Version: "1.0", Data: []udata{links, udata{ID: "tasks", Data: []udata{}}}, Error: []udata{}}}
}

func mkError(name, rel, value string) []byte {
	bs, err := json.Marshal(udoc{ubody{Version: "1.0", Error: []udata{udata{Name: name, Rel: []string{rel}, Value: value}}}})
	if err != nil {
		panic(err)
	}
	return bs
}
