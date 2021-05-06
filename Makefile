#!/usr/bin/make -f
# ----------------------------------------------------------------------
#

build:
	go build .

mock-install:
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen

mock-gen:
	mockgen github.com/t4rest/t4rest-service/redis MyCacher > ./redis/mock_redis/redis.go
	mockgen github.com/t4rest/t4rest-service/dao DAO > ./dao/mock_dao/dao.go
