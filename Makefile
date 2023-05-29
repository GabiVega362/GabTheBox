.PHONY: default
default: run

.PHONY: run
run: main.go
	@docker-compose up -d
	@go run main.go

.PHONY: deploy
deploy: docker-compose.yml
	@docker-compose up -d
	@docker-compose logs -f

.PHONY: connect
connect: docker-compose.yml
	@docker-compose exec database psql -U gabthebox

.PHONY: clean
clean:
	@rm -rf ./dist

.PHONY: purge
purge: docker-compose.yml clean
	@docker-compose down -v --remove-orphans --rmi all