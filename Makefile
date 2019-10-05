CONTAINER_PORT?=6000
PUBLIC_PORT?=9000
TAG?=$(shell echo "prod-$(shell git rev-list HEAD --max-count=1 --abbrev-commit)")
export TAG
export PORT
export PUBLIC_PORT

test:
	go test ./...

pack:
	docker build --build-arg main-version-tag="-X main.version=$(TAG)" --build-arg port=$(PORT) -t xman2019/sum-microservice:$(TAG) -f Dockerfile .

upload:
	docker push xman2019/sum-microservice:$(TAG)

deploy:
	envsubst < k8s-deployment.yml | kubectl apply -f -

ship: test pack upload deploy
