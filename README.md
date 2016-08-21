## API Resource Server

[![Build Status](https://travis-ci.org/tquach/golang-rest-server.png?branch=master)](https://travis-ci.org/tquach/golang-rest-server) [![Coverage Status](https://coveralls.io/repos/tquach/golang-rest-server/badge.png?branch=master)](https://coveralls.io/r/tquach/golang-rest-server?branch=master) [![Stories in Ready](https://badge.waffle.io/tquach/golang-rest-server.png?label=ready&title=Ready)](https://waffle.io/tquach/golang-rest-server)

This is an example of a basic API server that retrieves resources based on the URL pattern. It relies on a MongoDB instance and standard REST conventions.

## Getting Started

Seed your MongoDB with some data. In your mongo client, run the following:

    use test
    db.notes.insert({"note": "This is a note."});

## Start up the application server

1. Install Godep with `go get -u github.com/tools/godeps`.
2. Build the application with `make build`.
3. Run the application within Docker or standalone.

## Run modes
* With Docker
	* Run `make start`
	* Access the server at `http://$DOCKER_IP:9000/`
* Standalone
	* Run `./golang-rest-server` will start the server at `http://localhost:9000/` and look for a mongo DB at `localhost`.
	* Use command line arguments to override: `./golang-rest-server --help`