FROM golang:1.15-alpine
RUN apk update && apk add git

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/gorilla/mux && go get github.com/go-sql-driver/mysql
RUN go build -o main .

EXPOSE 10000

CMD ["/app/main"]