#!/bin/sh

docker build --file Dockerfile --tag water-jug-riddle .
go test ./...
docker run -it water-jug-riddle