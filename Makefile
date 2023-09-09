init:
	@touch main.go
	@echo "package main" > main.go
	@touch input.txt

run:
	@go run main.go < input.txt
