postgres-setup:
	docker-compose up -d

postgres-destroy:
	docker-compose stop
	docker-compose rm -f
	rm -rf ~/gateway-router/postgres

createdb:
	docker exec -it db-db-1  createdb --username=root --owner=root gateway_router

dropdb:
	docker exec -it db-db-1 dropdb  gateway_router

migrate-up:
	migrate -path migration -database "postgresql://root:admin@localhost:5453/gateway_router?sslmode=disable" -verbose up

migrate-down:
	migrate -path migration -database "postgresql://root:admin@localhost:5453/gateway_router?sslmode=disable" -verbose down 

refresh:
	docker exec -it db-db-1 dropdb  gateway_router
	docker exec -it db-db-1  createdb --username=root --owner=root gateway_router
	migrate -path migration -database "postgresql://root:admin@localhost:5453/gateway_router?sslmode=disable" -verbose up
	sqlc generate 

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres-destroy postgres-setup migrate-up migrate-down sqlc test refresh