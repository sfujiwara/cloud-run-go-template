.PHONY: hello
hello:
	echo hello

.PHONY: docker-build
docker-build:
	go build -o _build/server
	docker build -t cloud-run-go-template .

.PHONY: docker-run
docker-run:
	docker run --rm -it cloud-run-go-template

.PHONY: deploy
deploy:
	gcloud run deploy cloud-run-go-template \
	  --project sfujiwara-dev \
	  --region us-central1 \
	  --allow-unauthenticated \
	  --image us-central1-docker.pkg.dev/sfujiwara-dev/docker/cloud-run-go-template
