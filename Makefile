up:
	docker-compose up -d

be-run:
	docker-compose exec app go run cmd/srv/main.go

be-log:
	docker-compose logs --tail 100 -f app

restart:
	docker-compose restart

go-lint:
	docker-compose exec app golangci-lint run

go-lint-fmt:
	docker-compose exec app gofmt -w app pkg sandbox testsupport cmd

migrate-new:
	docker-compose exec migrate sql-migrations new ${name}

migrate-%:
	$(eval CMD:= $*)
	docker-compose run migrate sql-migrate $(CMD) -env=development -config=dbconfig.yml; \
	seq 3 | xargs -P8 -I{} docker-compose run -e DBNAME="go-random_test{}" migrate sql-migrate $(CMD) -config=dbconfig.yml -env=test;

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
	docker-compose exec db mysql -urandom -prandom

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
