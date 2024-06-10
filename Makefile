build:
	@go build -o myapp

run: build
	@./myapp

scar: 
	@go run main.go