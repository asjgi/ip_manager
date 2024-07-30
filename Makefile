ECR_HOST := 123456789.dkr.ecr.ap-northeast-2.amazonaws.com
ifdef TAG
	TAG := $(TAG)
endif

ifndef TAG
    TAG := "latest"
 endif

IMAGE := ${ECR_HOST}/cluster/ip-manager:${TAG}

.PHONY: default build push

default: build

build:
	docker buildx build \
    --load \
    --platform=linux/arm64 \
    --tag ${IMAGE} \
    .

push: build
	docker push ${IMAGE}
