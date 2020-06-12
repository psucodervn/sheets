#!/usr/bin/env bash

TAG="v$(date -u +"%Y%m%d")-$(git rev-parse --short HEAD)"
export TAG

export SERVICE=$1

REPOSITORY=psucoder/sheets-${SERVICE}
export REPOSITORY

IMAGE=${REPOSITORY}:${TAG}
export IMAGE

echo "Build ${IMAGE}";

docker-compose -f dc-build.yaml build ${SERVICE}
docker-compose -f dc-build.yaml push ${SERVICE}

./up.sh ${SERVICE} ${TAG}

exit 0
