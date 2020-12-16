# GO-REST

## Install package

`export GOPATH=~/go`

`go get github.com/gorilla/mux`

## Run

Start web server:
`go run main.go`

Access web server:
http://localhost:10000/

Endpoints:
- GET http://localhost:10000/dishes
- GET http://localhost:10000/dish/1
- POST http://localhost:10000/dish
- DELETE http://localhost:10000/dish/1