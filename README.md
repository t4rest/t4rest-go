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

''
To upgrades all the direct and indirect dependencies of your module.
```
go get -u -t ./...
```

style guide    - https://github.com/uber-go/guide/blob/master/style.md
better errors  - https://blog.mediocregopher.com/2021/03/20/a-simple-rule-for-better-errors.html
retry          - https://github.com/avast/retry-go