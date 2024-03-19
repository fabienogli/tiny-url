help:		## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: 		## build stack
	docker-compose -f docker-compose.yml build
up: 		## start stack
	docker-compose -f docker-compose.yml up -d
down: 		## stop all stack
	docker-compose -f docker-compose.yml down
restart: 	## restart stack
	docker-compose -f docker-compose.yml stop
logs: 		## display logs
	docker-compose -f docker-compose.yml logs -f