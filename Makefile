build:
	@go build -o /bin/main ./main.go
run: build
	@/bin/main
generate:
	templ generate