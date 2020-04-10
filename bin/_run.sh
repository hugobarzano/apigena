#!/usr/bin/env bash

DOCKER_IMAGE="generative-env:latest"
#DOCKER_IMAGE="hugobarzano/dlps:latest"
DOCKER_IN_DOCKER="-v /var/run/docker.sock:/var/run/docker.sock"
ENTRYPOINT="/bin/bash"
WORKSPACE="/var/data/"
PORTS="-p 3000-3050:3000-3050"
RUN="docker run $PORTS --rm -v $(pwd):$WORKSPACE -it"

$RUN $DOCKER_IN_DOCKER --workdir $WORKSPACE $DOCKER_IMAGE $ENTRYPOINT
