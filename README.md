# golang-todo

### Overview

Simple Golang Client for SearchBlox API(v8.6) - https://www.searchblox.com/:
- Go(1.8.0) programming language - https://golang.org/
- Docker(17.03.0-ce) - https://github.com/docker/docker
- Docker-Compose(1.11.2) - https://github.com/docker/compose

To run Searchblox and golang application Git, Docker and Docker-Compose should be installed locally.
<br /> Docker install: https://docs.docker.com/engine/installation/
<br /> Docker-Compose install: https://docs.docker.com/compose/install/

### Usage
Open terminal and clone this repository by SSH(or HTTPS):
```
    git clone git@github.com:tokillamockingbird/searchblox-golang.git
```
Locate to searchblox-golang directory.
To start all needed services simply execute one command
```
    docker-compose up -d
```
After this you can check Searchblox on 8089 port, open `localhost:8089` in browser to check it.
