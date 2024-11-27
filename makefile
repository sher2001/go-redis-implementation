run: build
	@./bin/myRedis
build: 
	@go build -o bin/myRedis .
test : 
	@go test -v ./...