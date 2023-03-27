[![Go](https://github.com/ronnpang1/impactify-api/actions/workflows/go.yml/badge.svg)](https://github.com/ronnpang1/impactify-api/actions/workflows/go.yml)
## Document Purpose

The purpose of this document is mean't to discuss how to run this api locally.

## Run Locally
To run this project locally simply navigate to the root of this folder and run:

```console
$ docker-compose up -d --no-deps --build
```

This will create two running containers (the api and the mysql container). Swagger docs are available at: http://localhost:8080/swagger/index.html, but the following routes are also available on intialization:

 `publisher data rows route => [POST] http://localhost:8080/publisher/data/rows/:id/:currency`

 `publisher aggregate data route => [POST] http://localhost:8080/publisher/data/:id/:currency`

 `publisher retrieve all data route => [POST] http://localhost:8080/publisher/data/all/:currency`

 `get publisher => [GET] http://localhost:8080/publisher/:id`
 
 `get all publishers => [GET] http://localhost:8080/publishers`

If one wants to run the api locally, and having a running instance of the database, one would just need to do the following:
1) Change the config.toml and have have db point to localhost.
2) Run from root:
```console
$ docker-compose up -d --no-deps --build db
```
3) change config from db to localhost in ```config.toml```
    ```
    [db]
    host = "localhost"
    ```
4) add in api keys for fixer and exchange-rate api in ```config.toml```
    ```
    [api]
    fixer_api_key = "exampleapikey1"
    exchange_rate_api_key = "exampleapikey2"
    ```

5) Run from root (windows): 
```console
$ go run .\main.go 
```
or 

```console
$ go run ./main.go 
```