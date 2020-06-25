all:build
	docker-compose build --no-cache && docker-compose up

build: build_go build_posgre


build_go:
	docker image build ./ -t shin0501/devi-web:latest

go:
	docker container run -itd --rm -p 8080:8080 shin0501/devi-web:latest


build_posgre:
	docker image build ./docker/postgresql/ -t shin0501/devi-db:latest

posgre:
	docker container run --env-file ${pwd}/docker/.env -itd --rm -p 5432:5432 shin0501/devi-db:latest


up:
	docker-compose build --no-cache && docker-compose up

