run:
	go run main.go

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