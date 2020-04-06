SHELL            = /bin/bash

APP_NAME         = gohello
VERSION         := $(shell git describe --always --tags)
GIT_COMMIT       = $(shell git rev-parse HEAD)
GIT_DIRTY        = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true))
BUILD_DATE       = $(shell date '+%Y-%m-%d-%H:%M:%S')
DOCKER_REGISTRY  = zackijack
TEAM             = backend

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Make commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build            Compile the project.'
	@echo '    make package          Build final Docker image with just the Go binary inside.'
	@echo '    make tag              Tag image created by package with latest, Git commit and version.'
	@echo '    make push             Push tagged images to registry.'
	@echo '    make run ARGS=        Run with supplied arguments.'
	@echo '    make test             Run tests on a compiled project.'
	@echo '    make clean            Clean the directory tree.'

	@echo

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -X github.com/zackijack/gohello/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/zackijack/gohello/version.Version=${VERSION} -X github.com/zackijack/gohello/version.Environment=${ENVIRONMENT} -X github.com/zackijack/gohello/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

.PHONY: package
package:
	@echo "Building image ${APP_NAME} ${VERSION} ${GIT_COMMIT}"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=${GIT_COMMIT}${GIT_DIRTY} -t ${DOCKER_REGISTRY}/${APP_NAME}:local .

.PHONY: tag
tag: package
	@echo "Tagging: latest ${VERSION} ${GIT_COMMIT}"
	docker tag ${DOCKER_REGISTRY}/${APP_NAME}:local ${DOCKER_REGISTRY}/${APP_NAME}:${GIT_COMMIT}
	docker tag ${DOCKER_REGISTRY}/${APP_NAME}:local ${DOCKER_REGISTRY}/${APP_NAME}:${VERSION}
	docker tag ${DOCKER_REGISTRY}/${APP_NAME}:local ${DOCKER_REGISTRY}/${APP_NAME}:latest

.PHONY: push
push: tag
	@echo "Pushing Docker image to registry: latest ${VERSION} ${GIT_COMMIT}"
	docker push ${DOCKER_REGISTRY}/${APP_NAME}:${GIT_COMMIT}
	docker push ${DOCKER_REGISTRY}/${APP_NAME}:${VERSION}
	docker push ${DOCKER_REGISTRY}/${APP_NAME}:latest

.PHONY: deploy
deploy:
	@echo "Deploying ${APP_NAME} ${VERSION}"
	helm upgrade ${APP_NAME} machtwatch/app --install \
		--namespace ${TEAM} \
		--values _infra/helm/${ENVIRONMENT}.yaml \
		--set meta.env=${ENVIRONMENT},meta.maintainer=${TEAM},meta.version=${VERSION},image.repository=${REGISTRY_URL}/${APP_NAME},image.tag=${VERSION}

.PHONY: run
run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: test
test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test ./...

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}
