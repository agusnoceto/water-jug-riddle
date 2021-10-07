#!/bin/sh

docker build --file Dockerfile --tag water-jug-riddle .
docker run -it water-jug-riddle