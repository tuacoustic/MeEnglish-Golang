#GO MAKES ME BETTER 
#AUTHOR - TUDINH
PROJECT_NAME 	     := "MeEnglish Go API"
MODULE 			     := $(shell go list -m)
GIT_BRANCH           := $(shell git symbolic-ref HEAD | sed -e "s/^refs\/heads\///")
DATE 			     ?= $(shell date +%FT%T%z)

# GO_VERSION := $(shell go -v)
VERSION := "1.0.0"

#PWD
PROJECT_PWD 	   := $(shell pwd)

#COLOR
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
RESET 		 := $(shell tput -Txterm sgr0)
ALERT        := $(YELLOW)$(DATE)$(RESET)

#Docker 
# DOCKER_SERVICE_NAME  := "me-english-registry:5000/go-api:dev"
DOCKER_SERVICE_NAME  := "erp-registry:5000/me-english-go:prod"
DOCKER_NAME 		 := "me-english"
DOCKER_REPLICAS      := 1
DOCKER_PORT   		 := 4040:4040
DOCKER_NONE_IMAGES   := $(shell docker images --filter "dangling=true" -q --no-trunc)
DOCKER_SOURCE_FILE   := ${PROJECT_PWD}/docker/go-service/Dockerfile

#See version
.PHONY: version
version:
	@echo "Copyright © 2021 MeEnglish | Golang BackEnd | Version: $(VERSION)"
	@git branch -v
	@echo "See more: https://github.com/MeEnglish/golang"

# Package Reload
.PHONY: get-air
get-air: 
	@echo "[$(ALERT)] - make get-air -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@go get -u github.com/cosmtrek/air

.PHONY: hot-reload
hot-reload:
	@echo "[$(ALERT)] - make hot-reload -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@./bin/air

.PHONY: ins-swagger
ins-swagger:
	@echo "[$(ALERT)] - make ins-swagger -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: gen-swagger
gen-swagger:
	@echo "[$(ALERT)] - make gen-swagger -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@./bin/swagger generate spec -o ./swagger.yaml --scan-models

#Docker Build
.PHONY: docker-build
docker-build:
	@echo "[$(ALERT)] - make docker-build -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@docker build -t $(DOCKER_SERVICE_NAME) -f ${DOCKER_SOURCE_FILE} .
	@echo "[$(ALERT)] - make docker-build Finished 👌🏽"

#Docker Push
.PHONY: docker-push-service
docker-push-service:
	@echo "[$(ALERT)] - make docker-push-service -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@docker push $(DOCKER_SERVICE_NAME)
	@echo "[$(ALERT)] - make docker-push-service Finished 👌🏽"

#Docker Run Loadbalancing Service
.PHONY: docker-run-service
docker-run-service:
	@echo "[$(ALERT)] - make docker-run-service -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@docker service create --replicas $(DOCKER_REPLICAS) -p $(DOCKER_PORT) --name $(DOCKER_NAME) $(DOCKER_SERVICE_NAME) 
	@echo "[$(ALERT)] - make docker-run-service Finished 👌🏽"

.PHONY: docker-rm-service
docker-rm-service:
	@echo "[$(ALERT)] - make docker-rm-service -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@docker image prune -f
	@echo "[$(ALERT)] - make docker-rm-service Finished 👌🏽"

.PHONY: docker-update-service
docker-update-service:
	@echo "[$(ALERT)] - make docker-update-service -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@docker service update --image $(DOCKER_SERVICE_NAME) $(DOCKER_NAME)
	@echo "[$(ALERT)] - make docker-update-service Finished 👌🏽"