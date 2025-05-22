build:
	docker-compose build
run:
	docker-compose up -d
down:
	docker-compose down
info:
	docker ps -a

delete-client:
	docker rmi dips-client

restart-client:
	docker-compose down; \
	docker rmi dips-client; \
	sudo docker-compose up -d
