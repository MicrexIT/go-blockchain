# go-blockchain

This repo contains everything you need to get started building a blockchain in golang

## Requirements

To run this project, we advise to use Docker and Docker Compose.
It will avoid the trouble of installing all the required dependencies.
This projects uses Golang v1.4

## Quick start

The Makefile offers an extra layer on top of Docker Compose to simplify running the project.

Starting the project:
```bash
make start
``` 

Restarting the project and rebuild images if necessary
```bash
make restart
```

Display a list of the running services
```bash
make status
```

Stop the project and clean up environment
```bash
make stop
```

Inspect logs of the service selected with the variable `SERVICE`
```bash
make inspect SERVICE=blockchain
```


