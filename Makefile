APP_NAME=scheduler
DOCKER_COMPOSE=docker-compose


build:
	@echo "Building the project..."
	go build -o $(APP_NAME) ./cmd/api/main.go

run:
run:
	@echo "Running the application..."
	./scheduler

docker-up:
	@echo "Starting Docker containers..."
	$(DOCKER_COMPOSE) up --build -d

docker-down:
	@echo "Stopping Docker containers..."
	$(DOCKER_COMPOSE) down

clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(APP_NAME)

restart:
	@echo "Restarting Docker containers..."
	make docker-down
	make docker-up

help:
	@echo "Available commands:"
	@echo "  build        Build the Go application"
	@echo "  run          Run the Go application"
	@echo "  docker-up    Start Docker containers"
	@echo "  docker-down  Stop Docker containers"
	@echo "  clean        Clean build artifacts"
	@echo "  restart      Restart Docker containers"
