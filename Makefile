.PHONY: build up down logs restart

build:
	docker compose build

up:
	docker compose up

down:
	docker compose down

logs:
	docker compose logs -f --tail=200

restart:
	docker compose down
	docker compose build
	docker compose up -d
