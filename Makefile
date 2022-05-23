postgres:
	docker run --name my-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it my-postgres createdb --username=root --owner=root gmeet_db

dropdb:
	docker exec -it my-postgres dropdb gmeet_db

run:
	go run ./cmd/biz/main.go

schema:
	@read -p "Enter the schema name: " name; \
		go run entgo.io/ent/cmd/ent init $$name

generate:
	go generate ./ent

mac:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/biz

linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/biz

windows:
	env GOOS=windows GOARCH=amd64 go build -o ./bin/windows/biz.exe