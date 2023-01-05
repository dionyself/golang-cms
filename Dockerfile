FROM golang:1.19.4-alpine3.17
LABEL "email"="dionyself@gmail.com"

RUN apk add --no-cache git gcc g++
RUN mkdir /app
RUN cd /app && git clone https://github.com/dionyself/golang-cms && cd golang-cms \
 && go get -u github.com/beego/bee \
 && go mod tidy

EXPOSE 8080

# Start app
CMD sh /app/golang-cms/run.sh
