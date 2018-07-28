NAME = procfile-util
MAINTAINER = josegonzalez
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.1
IMAGE_NAME ?= $(NAME)
BUILD_TAG ?= dev

build:
	@$(MAKE) build/darwin/$(NAME)
	@$(MAKE) build/linux/$(NAME)
	@$(MAKE) docker-image

build/darwin/$(NAME):
	mkdir -p build/darwin && CGO_ENABLED=0 GOOS=darwin go build -a -ldflags "-X main.Version=$(VERSION)" -o build/darwin/$(NAME)

build/linux/$(NAME):
	mkdir -p build/linux  && CGO_ENABLED=0 GOOS=linux  go build -a -ldflags "-X main.Version=$(VERSION)" -o build/linux/$(NAME)

clean:
	rm -rf build

circleci:
	docker version
	rm -f ~/.gitconfig

deps:
	go get -u github.com/progrium/gh-release/...
	dep ensure -vendor-only

docker-image:
	docker build -q -f Dockerfile.build -t $(IMAGE_NAME):$(BUILD_TAG) .

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(HARDWARE).tgz -C build/linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz -C build/darwin $(NAME)
	gh-release create $(MAINTAINER)/$(NAME) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)
