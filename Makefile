build-image:
	docker build -t rafaelcmedrado/desafio:latest -f build/Dockerfile .

push-image: build-image
	docker push rafaelcmedrado/desafio:latest

run-local:
	docker-compose -f deploy/local/docker-compose.yml up