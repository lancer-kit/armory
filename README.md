[![GoDoc](https://godoc.org/github.com/lancer-kit/armory?status.png)](https://godoc.org/github.com/lancer-kit/armory)
[![Go Report Card](https://goreportcard.com/badge/github.com/lancer-kit/armory)](https://goreportcard.com/report/github.com/lancer-kit/armory)

# Armory Service Kit

Common libraries for building go services:

## Install

```shell script
go get -u github.com/go-chi/chi
```

## Features 

- **Api**
    - [Render](./api/render/README.md) - response helper, base responses
    - [HTTPX](./api/httpx) - wrapper for `http.Client` with additional helpers the for RESTfull APIs. 
- [DB](./db/README.md) - connector for the ORMless interaction with the PostgreSQL databases.
- [Log](./log/README.md) - simple wrapper for logrus with some useful perks.

- [Auth](./auth/README.md) - methods for the service authorization.
- [Crypto](./crypto/README.md) - wrappers for hashing, signing, random values generation etc.
- [natsx](./natsx/README.md) - simple wrapper for NATS.


## Usage 

For details and documentation please check the [godoc.org](https://godoc.org/github.com/lancer-kit/armory) 
 
- [api/httpx](https://godoc.org/github.com/lancer-kit/armory/api/httpx)
- [api/render](https://godoc.org/github.com/lancer-kit/armory/api/render)
- [auth](https://godoc.org/github.com/lancer-kit/armory/auth)
- [crypto](https://godoc.org/github.com/lancer-kit/armory/crypto)
- [db](https://godoc.org/github.com/lancer-kit/armory/db)
- [db/test](https://godoc.org/github.com/lancer-kit/armory/db/test)
- [initialization](https://godoc.org/github.com/lancer-kit/armory/initialization)
- [log](https://godoc.org/github.com/lancer-kit/armory/log)
- [natsx](https://godoc.org/github.com/lancer-kit/armory/natsx)
- [tools](https://godoc.org/github.com/lancer-kit/armory/tools)
- [tools/queue](https://godoc.org/github.com/lancer-kit/armory/tools/queue)

## Tools

- [Forge](https://github.com/lancer-kit/forge) — a common tool for code-generation and projects bootstrap, includes templates oriented for **Armory** usage.

## Examples

- [service-scaffold](https://github.com/lancer-kit/service-scaffold) - an example project built with **Armory**;
- [domain-based-scaffold](https://github.com/lancer-kit/service-scaffold) - other example project built with **Armory**;  


## Licence

Lancer-Kit Armory is Apache 2.0 licensed.

