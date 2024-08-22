.DEFAULT_GOAL := help

NAME=$(shell basename "$(PWD)")

# ##############################################################################
# # GENERAL
# ##############################################################################
.PHONY: help
help: makefile
	@echo
	@echo " Available actions in "$(NAME)""
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':'
	@echo

# ##############################################################################
# # RECIPES
# ##############################################################################
##: --- GO ---
##: run: go run
run:
	go run ./server/.

## wire: go di wire
wire:
ifeq ($(shell which wire),)
	go install github.com/google/wire/cmd/wire@latest
	cd di; \
	wire
else
	cd di; \
	wire
endif
	perl -0777 -i -pe 's/func\(\)\s*\{.*?cleanup\(\).*?\}\s*/cleanup\n/gs' di/wire_gen.go

## mock: go mock generator
mock:
ifeq ($(shell which mockery),)
	go install github.com/vektra/mockery/v2@latest
	cd server; \
	mockery --all --keeptree
else
	cd server; \
	mockery --all --keeptree
endif

## swagger: go swagger init
swagger:
ifeq ($(shell which swag),)
	go install github.com/swaggo/swag/cmd/swag@latest
	cd server; \
	swag init
	swag fmt
else
	cd server; \
	swag init
	swag fmt
endif

## m: go wire && go mock
m: wire mock

## :
# ##############################################################################
# # DOCKER
# ##############################################################################
##: --- Docker ---
## restart: docker down & up
restart: down up

## down: docker down
down:
	docker-compose -f docker-compose-test.yml down

## up: docker up
up:
	docker-compose -f docker-compose-test.yml up -d

## prune: docker prune
.PHONY: prune
prune:
	docker volume prune -f
	docker image prune -f
	docker container prune -f
	docker network prune -f
## :
# ##############################################################################
# # TEST
# ##############################################################################
##: --- Unit Test ---
## unit: run unit test (arg=<filename>)
unit:
ifeq ($(arg),)
	cd server; \
	go test ./...
else
	cd server; \
	go test $(shell find $(PWD) -name "$(arg)_test.go") $(shell find $(PWD) -name "$(arg).go")
endif
## :
##: --- Acceptance Test ---
## test: run acceptance test (arg=<filename>)
test:
ifeq ($(arg),)
	cd acceptance_test; \
	npm run cy:run
else
	cd acceptance_test; \
	npm run cy:run -- --spec $(shell find $(PWD) -name "$(arg).cy.js" | sed 's/.*acceptance_test\///')
endif