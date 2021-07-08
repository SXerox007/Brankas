# Brankas

## Overview:

DB: `postgres`

Server: `go`

Request Multiplexer: `gorilla/mux`

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

https://github.com/gorilla/mux

And Docker containers used.


## How to Run

1. Clone the `repo`

2. run through `docker,makefile,or run.sh file`

```
sh deployment/run.sh  (it contains the docker commands)
```

Before running `make app` export all env (change `PG_HOST` to localhost and `PG_HOST_WITH_PORT` to localhost:5432) i.e in deployment/local.env then `go build application.go middleware.go brankas.go`

```
make app 
```

3. Run the migrations i.e in base/db/postgres/sql (Manual as of now)

4. open browser http://localhost:50051/brankas/upload 



## Explaination (Short):
Code starts at `application.go`. First start the db and create mux for api request and start the server.

when you hit the http://localhost:50051/brankas/upload it will open a basic html page and asking you for image file to upload in UI.
In Backend what happens when you hit `/brankas/upload` it sends the `auth` token into that page.

When FE hit multipart request at `api/v1/brankas/upload` it will first go to the middleware there contains the check auth token is there in a request or not. If not it sends you error from there only.

After that it came to main `UploadFile` func 
there many checks 

1. get file check
2. file Size check Not exceed 8MB
3. Content-Type check

After that it create a new file in local and store image details in db under `all_images` table.


## Need Help:
sumitthakur769@gmail.com





