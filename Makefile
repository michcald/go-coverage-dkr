SHELL:=/bin/bash

export IMAGE_NAME:=michcald/go-coverage

default: login build push logout

login:
	@docker login --username ${DOCKERHUB_USERNAME} --password ${DOCKERHUB_PASSWORD}

build:
	docker build --rm --no-cache -t ${IMAGE_NAME}:${RELEASE_VERSION} .
	if [[ ${RELEASE_VERSION} =~ ^v[0-9]+.[0-9]+.[0-9]+$$ ]]; then \
		docker tag ${IMAGE_NAME}:${RELEASE_VERSION} ${IMAGE_NAME}:latest; \
	fi

push:
	docker push ${IMAGE_NAME}:${RELEASE_VERSION}
	if [[ ${RELEASE_VERSION} =~ ^v[0-9]+.[0-9]+.[0-9]+$$ ]]; then \
		docker push ${IMAGE_NAME}:latest; \
	fi

logout:
	docker logout
