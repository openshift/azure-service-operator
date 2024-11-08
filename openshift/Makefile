# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/..)


# Define Docker related variables. Releases should modify and double check these vars.
ifeq (,$(shell command -v gcloud))
    REGISTRY ?= gcr.io/$(shell gcloud config get-value project)
else
    REGISTRY ?= localhost:5000
endif
IMG_NAME ?= aso-controller
CONTROLLER_IMG ?=$(REGISTRY)/$(IMG_NAME)
ARCH := amd64
TAG := latest


.PHONY: binaries
binaries:
	go build -C ${ROOT_DIR}/v2 -o aso-controller ./cmd/controller

# Redefine this target since upstream uses make for v1 and task for v2.
# NOTE: We use the repo root as the docker build context, but the Dockerfile itself
# will pull in files from the v2 directory
.PHONY: docker-build
docker-build:
	docker build -f openshift/Dockerfile --ignorefile openshift/Dockerfile.dockerignore -t $(CONTROLLER_IMG)-$(ARCH):$(TAG) ${ROOT_DIR}

.PHONY: docker-push
docker-push:
	docker push $(CONTROLLER_IMG)-$(ARCH):$(TAG)

.PHONY: test-unit
test-unit:
	${ROOT_DIR}/openshift/unit-tests.sh
