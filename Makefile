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
	mockgen github.com/t4rest/t4rest-go/mysql Tx > ./mysql/mock_tx/mytx.go
