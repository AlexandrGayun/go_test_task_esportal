include .env
DBCON_STRING = "${DB_USER}:${DB_PASSWORD}@(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true"
TEST_DBCON_STRING = "${TEST_DB_USER}:${TEST_DB_PASSWORD}@(${TEST_DB_HOST}:${TEST_DB_PORT})/${TEST_DB_NAME}?parseTime=true"

migrateup:
	@docker exec -it go_api_container migrate -source file://db/migrations/ -database mysql://${DBCON_STRING} up
migratedown:
	@docker exec -it go_api_container migrate -source file://db/migrations/ -database mysql://${DBCON_STRING} down
populateseeds:
	@docker exec -it go_api_container goose -no-versioning -dir "db/seeds/" mysql ${DBCON_STRING} up
migrateuptest:
	@docker exec -it go_api_container migrate -source file://db/migrations/ -database mysql://${TEST_DBCON_STRING} up
migratedowntest:
	@docker exec -it go_api_container migrate -source file://db/migrations/ -database mysql://${TEST_DBCON_STRING} down
populateseedstest:
	@docker exec -it go_api_container goose -no-versioning -dir "db/seeds/" mysql ${TEST_DBCON_STRING} up
test:
	docker exec -it go_api_container go test -v ./...
autotest: migrateuptest populateseedstest test migratedowntest
benchmark:
	docker exec -it go_api_container go test ./... -bench=. -benchmem
.PHONY: migrateup migratedown populateseeds migrateuptest migratedowntest populateseedstest test benchmark cleartest