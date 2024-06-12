
# Harmoniq API v2

This API for Harmoniq App using Golang 1.22.4


## Requirement Installation

- [Docker](https://docs.docker.com/engine/install/) version 25.0.1
- [docker-compose](https://docs.docker.com/compose/) latest version
- [Postman](https://www.postman.com/) latest version
- [Golang](https://go.dev/doc/install) version 1.22.4


## Serve
    make up = create container and image in docker, so it can be serve a server and running the application
    make down = shut down image and container in docker which is running by system
    make destroy = delete all container and image in docker
    make log = watch your development while it's running
## Endpoint Testing
    available in `./doc/Harmoniq.postman_collection.json` and ready to import to postman
