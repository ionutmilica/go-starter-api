#!/bin/bash

function stopAll {
    docker-compose -f ${composeFile} down
}

function build {
    docker-compose -f ${composeFile} build
}

function startDatabase {
    echo "Starting the database server..."
    docker-compose -f ${composeFile} up -d database

    echo "Waiting for the server to boot:"

    while ! mysqladmin ping -h'127.0.0.1' -u'root' -p'secret' --silent  > /dev/null 2>&1; do
        printf "."
        sleep 1
    done

    echo ""
}

function startCache {
    echo "Starting the cache server..."
    docker-compose -f ${composeFile} up -d cache
}

function startApi {
    echo "Starting the api server..."
    docker-compose -f ${composeFile} up -d api
}

function buildAndStartAll {
    stopAll
    build
    startDatabase
    startCache
    startApi
}

function buildAndStartDeps {
    stopAll
    startDatabase
    startCache
}

composeFile=artifacts/docker-compose.yml

case "$1" in
"start_all")
    buildAndStartAll
;;
"start_deps")
    buildAndStartDeps
;;
*)
    echo "Usage: ./dev.sh start_all | start_deps"
    exit 1
;;
esac
