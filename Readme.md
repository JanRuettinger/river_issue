# README

## Run locally

`make run/live`

## Run docker image

`sed 's/export //g' .envrc > .docker_env`
`docker build --tag yosemite .`
`docker run -p 4000:4000 --env-file ./.docker_env yosemite`

## Installation fixes

Find cairo installation folder: `pkg-config --cflags  -- cairo cairo`
Copy cairo folder into <user>/go/pkg/mod/github.com/ungerik/go-cairo@v0.0.0-20210317133935-984b32e6bac6
