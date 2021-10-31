ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	if [ `docker network ls | grep shared-local |wc -l` -eq 0 ]; then docker network create shared-local; fi
	$(ENV_LOCAL) docker-compose up


.PHONY: generate
generate:
	rm -rf mock/
	go generate ./...

.PHONY: test
test:
	$(ENV_LOCAL) go test ./test/... -count=1

.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --enable=gosec,prealloc,gocognit
