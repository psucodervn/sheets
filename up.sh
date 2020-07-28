#!/usr/bin/env bash

REMOTE=ec2-user@ec2-3-0-57-42.ap-southeast-1.compute.amazonaws.com
REMOTE_DIR=/home/ec2-user/sheet
SERVICE=${1}
TAG=${2}

if [[ -n ${SERVICE} ]]; then
  sed -i.bak "s/-${SERVICE}:.*/-${SERVICE}:${TAG}/g" deploy/docker-compose.yaml
fi

rsync -a ./deploy/ ${REMOTE}:${REMOTE_DIR}

ssh ${REMOTE} "cd $REMOTE_DIR && ./up.sh"
