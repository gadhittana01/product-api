# Product API
This repository is an API where you can do a CRUD operations.

# Pre-requisite to contribute
To run the service on your local machine, you must have these components installed on your local machine : 
- Golang compiler

# Getting started
- To run this project you can use this command below in the root of the repository:
```sh
// install the dependency
$ go get

// create a database in postgres with this specification
$ name : productDB
$ user : postgres
$ pass : 

// start a table migration 
$ make db-init

// start a seeder (optional)
$ make run-seed

// start the redis
$ make run-redis

// start the http server
$ make run-http-server-local 
```

# API Documentation
for the API documentation you can find it in the root of the folder with file name :
```sh
$ api-documentation.txt
```
