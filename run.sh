#!/bin/bash

cd /app/golang-cms
git pull origin master --tags
export GOLANG_CMS_VERSION=$(git tag --sort=committerdate | tail -1)
git checkout $GOLANG_CMS_VERSION
go get -u github.com/beego/bee/v2
go mod tidy
bee run
