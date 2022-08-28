DOCKER_COMPOSE_IMPL=docker-compose

.PHONY: dc/build
dc/build:
	$(DOCKER_COMPOSE_IMPL) build

.PHONY: dc/up
dc/up:
	$(DOCKER_COMPOSE_IMPL) up
