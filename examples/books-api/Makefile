# Target: Manage Docker services
services:
	@echo "Restarting Docker services..."
	docker compose down -v --remove-orphans
	docker compose up -d
	@echo "Docker services are up and running."

# Target: Initialize Go dependencies
init:
	@echo "Initializing Go modules..."
	go mod tidy
	go mod vendor
	@echo "Go modules initialized."

# Target: Run the application
run:
	@echo "Starting the application..."
	go run main.go rest

# Target: Build the application
build:
	@echo "Building the application..."
	go build -o $(GOPATH)/bin/books-api .
	@echo "Build completed. Binary available at $(GOPATH)/bin/books-api."

# Target: Start everything (services, init, DB check, run)
start: services init check-db run

# Target: Check if the database is up and running
check-db:

# Declare these targets as phony (not associated with real files)
.PHONY: services init run build start check-db
