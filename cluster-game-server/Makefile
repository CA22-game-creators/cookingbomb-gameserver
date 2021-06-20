ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	$(ENV_LOCAL) docker-compose up

.PHONY: build
build:
	docker-compose build

.PHONY: build-run
build-run: build run

.PHONY: build-env
build-env:
	if [ `docker network ls | grep shared-local |wc -l` -eq 0 ]; then docker network create shared-local; fi
