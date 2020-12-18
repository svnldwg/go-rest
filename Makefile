build:
	docker build -t go-rest .

run:
	docker run -p 10000:10000 -it go-rest

run-detached:
	docker run -p 10000:10000 -d go-rest

start:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down