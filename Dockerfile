FROM golang:1.10-alpine3.7
MAINTAINER Dionys Rosario <dionyself@gmail.com>

RUN apk add --no-cache git gcc g++ \
 && go get -u github.com/beego/bee \
 && go get -u github.com/dionyself/golang-cms

EXPOSE 8080

# Start app
CMD sh $GOPATH/src/github.com/dionyself/golang-cms/run.sh
