include config.mk

NAME := cloud-run-go-template
DOCKER_IMAGE := $(LOCATION)-docker.pkg.dev/$(PROJECT)/docker/$(NAME)


.PHONY: go-run
go-run:
	go run ./src

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE)

.PHONY: cloud-build
cloud-build:
	gcloud builds submit --project $(PROJECT) --tag $(DOCKER_IMAGE)

.PHONY: cloud-run
cloud-run:
	gcloud run deploy $(NAME) \
	  --project $(PROJECT) \
	  --region $(LOCATION) \
	  --image $(DOCKER_IMAGE) \
	  --max-instances 10 \
	  --port 80
