THIS_FILE := $(lastword $(MAKEFILE_LIST))

build:
	docker-compose -f docker-compose.yml build
up:
	docker-compose -f docker-compose.yml up -d
down:
	docker-compose -f docker-compose.yml down
restart:
	docker-compose -f docker-compose.yml stop
logs:
	docker-compose -f docker-compose.yml logs -f