#!/bin/bash

export GOLANG_CMS_VERSION=${GOLANG_CMS_VERSION:="master"}
cd $GOPATH/src/github.com/dionyself/golang-cms
bee run

git fetch $GOLANG_CMS_VERSION
git checkout $GOLANG_CMS_VERSION
git pull origin $GOLANG_CMS_VERSION
bee run
