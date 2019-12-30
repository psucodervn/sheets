#!/usr/bin/env bash

TAG=$(git describe --abbrev=0 --tags)
if [[ -z $TAG ]]; then
  TAG=$(git rev-parse --short HEAD)
fi
export TAG

export SERVICE=$1

# case $2 in
#   dev)
#     REPOSITORY=gcr.io/polaris-240607;;
#   pro)
#     REPOSITORY=gcr.io/polaris-production-v1;;
#   aws)
#     REPOSITORY=331585694431.dkr.ecr.ap-southeast-1.amazonaws.com/polaris;;
#   *)
#     REPOSITORY=gcr.io/$(gcloud config get-value project 2> /dev/null);;
# esac
REPOSITORY=psucoder/sheets-${SERVICE}
export REPOSITORY

IMAGE=${REPOSITORY}:${TAG}
export IMAGE

echo "Build ${IMAGE}";

docker-compose -f dc-build.yaml build $SERVICE
docker-compose -f dc-build.yaml push $SERVICE
exit 0
