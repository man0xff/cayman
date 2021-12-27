.PHONY: PHONY
PHONY:

mod_update: PHONY
	go mod vendor
	go mod tidy

mod_add: m ?=
mod_add: PHONY
	go get $(m)

run: PHONY
	go run ./cmd ~/my/cayman.conf.yaml

test: PHONY
	go test -v ./internal/core

env: PHONY
	docker 
