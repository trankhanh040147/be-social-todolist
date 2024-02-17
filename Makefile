include app.env
export

upv:
	docker compose up --build

up:
	docker compose up --build --detach

down:
	docker compose down

cache:
	docker build -t social-todo-service-cache -f Dockerfile-cache .

upcache:
	sed -i '' 's/Dockerfile/Dockerfile-with-cache/g' ./docker-compose.yaml
	docker compose up --build --detach
	sed -i '' 's/Dockerfile-with-cache/Dockerfile/g' ./docker-compose.yaml

network:
	docker network create $(NETWORK_NAME)

mysql:
	docker run --name $(DB_CONTAINER_NAME) -p $(DB_PORT):3306 -e MYSQL_ROOT_PASSWORD=$(DB_PASSWORD) -d $(MYSQL_IMAGE)

createdb:
	docker exec -it $(DB_CONTAINER_NAME) mysql -u$(DB_USERNAME) -p$(DB_PASSWORD) -e "CREATE DATABASE \`$(DB_NAME)\` DEFAULT CHARACTER SET = 'utf8mb4' DEFAULT COLLATE = 'utf8mb4_0900_ai_ci';"

dropdb:
	docker exec -it $(DB_CONTAINER_NAME) mysql -u$(DB_USERNAME) -p$(DB_PASSWORD) -e "DROP DATABASE \`$(DB_NAME)\`;"

db:
	docker exec -it $(DB_CONTAINER_NAME) mysql -u$(DB_USERNAME) -p$(DB_PASSWORD)

server:
	go run .

proto:
	rm -rf pb/*.go
	buf generate

evans:
	evans --host localhost --port 9099 -r repl

build:
	go build -o app

outenv:
	./app outenv

outenvfile:
	./app outenv > .env

.PHONY: upv up down cache upcache network mysql createdb dropdb db server proto evans build outenv outenvfile