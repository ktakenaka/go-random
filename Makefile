up:
	docker-compose up -d

be-run:
	docker-compose exec app go run cmd/srv/main.go

restart:
	docker-compose restart

go-lint:
	docker-compose exec app golangci-lint run

go-lint-fmt:
	docker-compose exec app gofmt -w app

time := $(shell date +%s)
create-migrate:
	@docker-compose exec app touch db/migrations/$(time)_$(name).up.sql
	@docker-compose exec app touch db/migrations/$(time)_$(name).down.sql

migrate-up:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations up

migrate-down:
	docker-compose exec app migrate -database mysql://random:random@tcp\(db:3306\)/go-random?multiStatements=true -path db/migrations down

test:
	docker-compose exec app go test ./...

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

be-bash:
	@docker-compose exec app bash

fe-sh:
	@docker-compose exec web sh

mysql:
	docker-compose exec db mysql -urandom -prandom go-random

yarn-install:
	docker-compose exec web yarn install

yarn-prettier-write:
	docker-compose exec web yarn prettier --write src

yarn-add:
	docker-compose exec web yarn add ${PKG}

yarn-lint:
	docker-compose exec web yarn lint

yarn-fmt:
	docker-compose exec web yarn lint-fmt
