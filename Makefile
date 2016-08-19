.PHONY: docker

EXECUTABLE ?= drone-rollbar
IMAGE ?= rschmukler/drone-rollbar

docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(EXECUTABLE)
	docker build --rm -t $(IMAGE) .
