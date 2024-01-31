.PHONY: up down prepare-test run-test

up:
	docker compose up --detach

down:
	docker compose down

prepare-test:
	git submodule update --init && make --directory=gonkey build

run-test:
	./gonkey/gonkey -host localhost:8080 -tests ./test.yaml

