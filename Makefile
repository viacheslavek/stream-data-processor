
.PHONY: build
build:
	docker-compose build

# TODO: узнать, как можно флагами это задавать

.PHONY: buildPostgres
buildPostgres:
	docker compose -f docker-compose.yml up -d postgres

.PHONY: buildRedis
buildRedis:
	docker compose -f docker-compose.yml up -d redis

.PHONY: buildCassandra
buildCassandra:
	docker compose -f docker-compose.yml up -d cassandra
