init:
	@if [ -f main.go ]; then \
		echo "main.go already exists. Exiting."; \
		exit 1; \
	fi
	@touch main.go
	@echo "package main\n\nfunc main() {\n}" > main.go
	@touch input.txt

run:
	@go run main.go < input.txt
