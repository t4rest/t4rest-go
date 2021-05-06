# Coding standards

To avoid hard times during merges, please use a conventional formating:

```
go fmt ./...
```

To ensure your code is as clean as it can be you can use `golangci-lint`:

```
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
  --disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
  --enable=structcheck --enable=errcheck --enable=dupl --enable=ineffassign \
  --enable=unconvert --enable=goconst --enable=gosec --enable=megacheck

```

Go mod
```
https://encore.dev/guide/go.mod
```


Resources:
```
A tour of Golang                 - https://tour.golang.org/
Annotated example programs       - https://gobyexample.com 
Golang snippets                  - https://www.golangprograms.com/
Golang tutorial series           - https://golangbot.com/learn-golang-series/
Golang course                    - https://tutorialedge.net/course/golang/
Resources for new Go programmers - https://dave.cheney.net/resources-for-new-go-programmers
Uber Go Style Guide              - https://github.com/uber-go/guide/blob/master/style.md
```

Libs:
```
List of usefull libraries - https://awesome-go.com/
Retry                     - https://github.com/avast/retry-go
```

Tools:
```
JSON to Golang struct     - https://mholt.github.io/json-to-go/
```

Tips:
```
Slice tricks       - https://github.com/golang/go/wiki/SliceTricks
Error programming  - https://peter.bourgon.org/blog/2019/09/11/programming-with-errors.html
Better errors      - https://blog.mediocregopher.com/2021/03/20/a-simple-rule-for-better-errors.html
Effective Go       - https://golang.org/doc/effective_go.html
https://peter.bourgon.org/blog/
```

Internals:
```
Slices internals       - https://blog.golang.org/go-slices-usage-and-internals
Golang modules         - https://blog.golang.org/using-go-modules
Scheduler saga         - https://www.youtube.com/watch?v=YHRO5WQGh0k
Understanding channels - https://www.youtube.com/watch?v=KBZlN0izeiY
```

Books:
```
An introduction to programming in Go  - https://www.golang-book.com/books/intro
Go bootcamp                           - http://www.golangbootcamp.com/book
```