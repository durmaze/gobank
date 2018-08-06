#!/bin/sh

export DOCKER_HOST="tcp://0.0.0.0:2375"
docker run -p 2525:2525 -p 6789:6789 expert360/mountebank start --mock
