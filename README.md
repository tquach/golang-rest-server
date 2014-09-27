## API Resource Server

[![Build Status](https://travis-ci.org/tquach/golang-rest-server.png?branch=master)](https://travis-ci.org/tquach/golang-rest-server) [![Coverage Status](https://coveralls.io/repos/tquach/golang-rest-server/badge.png?branch=master)](https://coveralls.io/r/tquach/golang-rest-server?branch=master)

This is an example of a basic API server that retrieves resources based on the URL pattern. It relies on a MongoDB instance and standard REST conventions.

## Getting Started

Seed your MongoDB with some data. In your mongo client, run the following:

    use test
    db.notes.insert({"note": "This is a note."});

1. Clone the repo.
2. `go get -d ./...` or if you have gpm installed `gpm`.
3. Run `go run server.go --databaseName=test --databaseUrl=localhost:27017`
4. Go to `http://localhost:3000/notes`

This is a work in progress.
