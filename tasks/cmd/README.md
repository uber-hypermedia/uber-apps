# taskd -- Tasks Hypermedia Server

_taskd_ serves a JSON version of the Uber media type that conforms to the _tasks profile_.

It is written in *go* and uses the Gorilla toolkit's  _mux_ and _handlers_ packages. To
build the server, assuming you have *go* installed and your _$GOPATH_ set up, type:

```
$ go install
```

To run it: 

```
$ $GOPATH/bin/taskd
```

Once it's installed you can try it out using curl, _e.g_, 

```
$ curl -X GET http://localhost:3006/tasks/
```
