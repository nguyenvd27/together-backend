DOCKER_BUILD_BASE := docker-compose -f docker-compose.yml

build-backend-server:
	docker build -t together_backend:latest -f ./deployments/Dockerfile .
	$(DOCKER_BUILD_BASE) build together_backend

start-backend-server:
	$(DOCKER_BUILD_BASE) up -d together_backend

stop-backend-server:
	$(DOCKER_BUILD_BASE) down
