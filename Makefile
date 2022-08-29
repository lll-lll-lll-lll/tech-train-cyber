.PHONY: dc/build
dc/build:
	docker-compose build

.PHONY: dc/up
dc/up:
	docker-compose up

.PHONY: dc/down-remove
dc/down-remove:
	docker-compose down --rmi all --volumes --remove-orphans