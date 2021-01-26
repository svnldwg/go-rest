start:
	docker-compose up -d

stop:
	docker-compose stop

build:
	docker-compose build

down:
	docker-compose down

logs:
	docker-compose logs -f go-rest