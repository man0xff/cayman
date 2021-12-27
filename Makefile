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

env: db PHONY

# db ===========================================================================

db_container = cayman_db
db_password = hjnahjyn
db_name = cayman

db: PHONY
	docker run -d -p 3306:3306 --name $(db_container) \
		-e MYSQL_DATABASE=$(db_name) \
		-e MYSQL_ROOT_PASSWORD=$(db_password) \
		mysql

db_enter: PHONY
	docker exec -it $(db_container) mysql -u root -p$(db_password) $(db_name)

db_clean: PHONY
	docker rm -f $(db_container)

# db ^==========================================================================
