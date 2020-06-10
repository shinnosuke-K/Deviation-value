all:build
	docker-compose build --no-cache && docker-compose up

build: go posgre

go:
	docker image build ./ -t shin0501/devi-web:latest

posgre:
	docker image build ./docker/postgresql/ -t shin0501/devi-db:latest

up:
	docker-compose build --no-cache && docker-compose up

