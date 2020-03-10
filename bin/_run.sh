#!/usr/bin/env bash
REPO=$(basename `git rev-parse --show-toplevel`)
docker run -p 6666:6666 $REPO
