FROM golang:1.14-alpine
RUN apk update && apk add git

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/gorilla/mux
RUN go build -o main .

EXPOSE 10000

CMD ["/app/main"]