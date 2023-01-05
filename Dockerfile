FROM golang:1.19.4-alpine3.17
LABEL "email"="dionyself@gmail.com"

RUN apk add --no-cache git gcc g++
RUN mkdir /app
RUN cd /app && git clone https://github.com/dionyself/golang-cms && sleep 3
RUN cd /app/golang-cms && git pull origin master --tags && git checkout $(git tag --sort=committerdate | tail -1) && sleep 3
RUN cd /app/golang-cms && GOMAXPROCS=1 go get github.com/beego/bee/v2 && sleep 3
RUN cd /app/golang-cms && GOMAXPROCS=1 go install github.com/beego/bee/v2 && sleep 3
RUN cd /app/golang-cms && GOMAXPROCS=1 go mod tidy

EXPOSE 8080

# Start app
CMD sh /app/golang-cms/run.sh
