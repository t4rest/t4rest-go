#!/usr/bin/env bash

go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

mockgen github.com/t4rest/t4rest-go/redis Cacher > ./redis/mock_redis/redis.go
mockgen github.com/t4rest/t4rest-go/mysql Tx > ./mysql/mock_tx/mytx.go
