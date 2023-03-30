#!/bin/bash

docker stop postgres --force
docker rm postgres --force

docker run \
 -p 5432:5432 \
 -e POSTGRES_PASSWORD=password \
 --name postgres \
 postgres