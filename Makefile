NAME = procfile-util
EMAIL = procfile-util@josediazgonzalez.com
MAINTAINER = josegonzalez
MAINTAINER_NAME = Jose Diaz-Gonzalez
REPOSITORY = go-procfile-util
HARDWARE = $(shell uname -m)
VERSION ?= 0.0.2
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
	@$(MAKE) build/deb/$(NAME)_$(VERSION)_amd64.deb
	@$(MAKE) build/rpm/$(NAME)-$(VERSION)-1.x86_64.rpm
	@$(MAKE) docker-image

build/darwin/$(NAME):
	mkdir -p build/darwin && CGO_ENABLED=0 GOOS=darwin go build -a -ldflags "-X main.Version=$(VERSION)" -o build/darwin/$(NAME)

build/linux/$(NAME):
	mkdir -p build/linux  && CGO_ENABLED=0 GOOS=linux  go build -a -ldflags "-X main.Version=$(VERSION)" -o build/linux/$(NAME)

build/deb/$(NAME)_$(VERSION)_amd64.deb: build/linux/$(NAME)
	mkdir -p build/deb && fpm \
		--architecture amd64 \
		--category utils \
		--chdir build/linux \
		--description "$$DESCRIPTION" \
		--input-type dir \
		--license 'MIT License' \
		--maintainer "$(MAINTAINER_NAME) <$(EMAIL)>" \
		--name procfile-util \
		--output-type deb \
		--package build/deb/$(NAME)_$(VERSION)_amd64.deb \
		--prefix /usr/local/bin \
		--url "https://github.com/$(MAINTAINER)/$(REPOSITORY)" \
		--version $(VERSION) \
		--verbose \
		$(NAME)

build/rpm/$(NAME)-$(VERSION)-1.x86_64.rpm: build/linux/$(NAME)
	mkdir -p build/rpm && fpm \
		--architecture x86_64 \
		--category utils \
		--chdir build/linux \
		--description "$$DESCRIPTION" \
		--input-type dir \
		--license 'MIT License' \
		--maintainer "$(MAINTAINER_NAME) <$(EMAIL)>" \
		--name procfile-util \
		--output-type rpm \
		--package build/rpm/$(NAME)-$(VERSION)-1.x86_64.rpm \
		--prefix /usr/local/bin \
		--rpm-os linux \
		--url "https://github.com/$(MAINTAINER)/$(REPOSITORY)" \
		--version $(VERSION) \
		--verbose \
		$(NAME)

build-validate:
	mkdir -p validation
	ls -lah build/deb
	ls -lah build/rpm
	dpkg-deb --info build/deb/$(NAME)_$(VERSION)_amd64.deb
	dpkg -c build/deb/$(NAME)_$(VERSION)_amd64.deb
	cd validation && ar -x ../build/deb/$(NAME)_$(VERSION)_amd64.deb
	cd validation && rpm2cpio ../build/rpm/$(NAME)-$(VERSION)-1.x86_64.rpm > $(NAME)-$(VERSION)-1.x86_64.cpio
	ls -lah validation

clean:
	rm -rf build control.tar.gz data.tar.gz debian-binary release validation

circleci:
	docker version
	rm -f ~/.gitconfig
	sudo apt install apt-transport-https build-essential curl gnupg2 rpm rsync rubygems-integration ruby-dev ruby -qy
	sudo gem install --no-ri --no-rdoc --quiet fpm

deps:
	go get -u github.com/progrium/gh-release/...
	dep ensure -vendor-only

docker-image:
	docker build -q -f Dockerfile.build -t $(IMAGE_NAME):$(BUILD_TAG) .

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(HARDWARE).tgz -C build/linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz -C build/darwin $(NAME)
	cp build/deb/$(NAME)_$(VERSION)_amd64.deb release/$(NAME)_$(VERSION)_amd64.deb
	cp build/rpm/$(NAME)-$(VERSION)-1.x86_64.rpm release/$(NAME)-$(VERSION)-1.x86_64.rpm
	gh-release create $(MAINTAINER)/$(REPOSITORY) $(VERSION) $(shell git rev-parse --abbrev-ref HEAD)

store-artifacts: build
	mkdir -p /tmp/artifacts
	cp -r build/* /tmp/artifacts
