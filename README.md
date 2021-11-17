# Cachapi

Cachapi is simple key-value memory store application you can use it with http requests. It also saves to store value periodically(per 30 minutes) to json file with "timestamp-data.json" name under /tmp folder.

PS: Value can be any type: bool, string, int, array, object..

## API DOC

You can reach api documentation via:

[Swagger](https://cachapi.netlify.app/)

## Installation

Clone the repository and run main.go it will serve the http server to `http://localhost:3000` or use docker file!

## Test Env

You can use `http://18.234.112.151/api/v1` address for test requests