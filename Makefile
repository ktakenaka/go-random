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

migrate-new:
	docker-compose exec app sql-migrations new ${name}

migrate-up:
	docker-compose exec app sql-migrate up -env=development -config=db/dbconfig.yml

migrate-down:
	docker-compose exec app sql-migrate down -env=development -config=db/dbconfig.yml

test:
	docker-compose exec app go test ./...

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

gen:
	docker-compose exec app go generate ./...

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
