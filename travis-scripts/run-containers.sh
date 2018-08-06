#!/bin/sh

docker run -d -p 2525:2525 -p 6789:6789 expert360/mountebank start --mock
