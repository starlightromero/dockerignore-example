compose = docker-compose -f

dev = docker-compose.dev.yml
test = docker-compose.test.yml
prod = docker-compose.yml


build-dev:
	${compose} ${dev} build

start-dev:
	${compose} ${dev} up

stop-dev:
	${compose} ${dev} down

build-test:
	${compose} ${test} build

start-test:
	${compose} ${test} up

stop-test:
	${compose} ${test} down

build-prod:
	${compose} ${prod} build

start-prod:
	${compose} ${prod} up

stop-prod:
	${compose} ${prod} down