run:
	docker-compose up -d
down:
	docker-compose down
info:
	docker ps -a

delete-services:
	docker rmi bonus_service; \
	docker rmi flight_service; \
	docker rmi ticket_service; \
	docker rmi auth_service; \
	docker rmi gateway_service; \
	docker rmi statistics_service;

delete-all:
	docker rmi bonus_service; \
	docker rmi flight_service; \
	docker rmi ticket_service; \
	docker rmi auth_service; \
	docker rmi gateway_service; \
	docker rmi statistics_service; \
	rm -rf pg_data_bonus; \
	rm -rf pg_data_flight; \
	rm -rf pg_data_ticket; \
	rm -rf pg_data_auth; \
	rm -rf pg_data_statistics;

restart:
	docker-compose down; \
	docker rmi gateway_service; \
	docker-compose up -d

logs:
	docker logs gateway_service

ruff:
	ruff check . --fix; \
	ruff format .
