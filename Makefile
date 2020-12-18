.PHONY:	build push

IMAGE = quay.io/fortnox/sidecar-proxy
VERSION = 1.0.1

build:
	CGO_ENABLED=0 GOOS=linux go build
	docker build --pull --rm -t $(IMAGE):$(VERSION) .

push:	
	docker push $(IMAGE):$(VERSION)

all:	build push

