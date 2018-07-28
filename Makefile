NAME = procfile-util
MAINTAINER = josegonzalez
REPOSITORY = go-procfile-util
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.1
IMAGE_NAME ?= $(MAINTAINER)/$(NAME)
BUILD_TAG ?= dev
define DESCRIPTION
Utility that allows users to interact with Procfile files
Procfiles may be specified on stdin or via a flag, but
must always be valid yaml.
endef

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
	@$(MAKE) release/$(NAME)_$(VERSION)_amd64.deb
	@$(MAKE) release/$(NAME)-$(VERSION)-1.x86_64.rpm
	gh-release create $(MAINTAINER)/$(REPOSITORY) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

release/$(NAME)_$(VERSION)_amd64.deb: build/linux/$(NAME)
	fpm \
		--architecture amd64 \
		--category utils \
		--description "$$DESCRIPTION" \
		--input-type dir \
		--license 'MIT License' \
		--maintainer "Jose Diaz-Gonzalez <procfile-util@josediazgonzalez.com>" \
		--name procfile-util \
		--output-type deb \
		--package release/$(NAME)_$(VERSION)_amd64.deb \
		--url "https://github.com/$(MAINTAINER)/$(REPOSITORY)" \
		--version $(VERSION) \
		build/linux/$(NAME)=/usr/local/bin/$(NAME)

release/$(NAME)-$(VERSION)-1.x86_64.rpm: build/linux/$(NAME)
	fpm \
		--architecture x86_64 \
		--category utils \
		--description "$$DESCRIPTION" \
		--input-type dir \
		--license 'MIT License' \
		--maintainer "Jose Diaz-Gonzalez <procfile-util@josediazgonzalez.com>" \
		--name procfile-util \
		--output-type rpm \
		--package release/$(NAME)-$(VERSION)-1.x86_64.rpm \
		--rpm-os linux \
		--url "https://github.com/$(MAINTAINER)/$(REPOSITORY)" \
		--version $(VERSION) \
		build/linux/$(NAME)=/usr/local/bin/$(NAME)
