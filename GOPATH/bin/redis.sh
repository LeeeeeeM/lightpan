#!/usr/bin/env bash

docker rm -f lightpan-redis
docker run  --name lightpan-redis --net=bridge --restart=always -p 6379:6379 -d redis
docker ps |grep yourfs