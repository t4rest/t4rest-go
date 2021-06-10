#!/usr/bin/make -f
# ----------------------------------------------------------------------
#

build:
	go build .

mock-install:
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen

mock-gen:
	mockgen github.com/t4rest/t4rest-go/redis Cacher > ./redis/mock_redis/redis.go
	mockgen github.com/t4rest/t4rest-go/mysql Tx > ./mysql/mock_mysql/mytx.go

fmt:
	go fmt ./...

lint-install:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint-run:
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m \
		--disable-all --enable=deadcode  --enable=gocyclo --enable=golint --enable=varcheck \
		--enable=structcheck --enable=errcheck --enable=dupl --enable=ineffassign \
		--enable=unconvert --enable=goconst --enable=gosec --enable=megacheck
