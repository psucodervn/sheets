#!/usr/bin/env bash

TAG=$(git describe --abbrev=0 --tags)
if [[ -z ${TAG} ]]; then
  TAG=$(git rev-parse --short HEAD)
fi
export TAG

export SERVICE=$1

REPOSITORY=psucoder/sheets-${SERVICE}
export REPOSITORY

IMAGE=${REPOSITORY}:${TAG}
export IMAGE

echo "Build ${IMAGE}";

docker-compose -f dc-build.yaml build ${SERVICE}
docker-compose -f dc-build.yaml push ${SERVICE}
exit 0
