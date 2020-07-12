[![GoDoc](https://godoc.org/github.com/lancer-kit/armory?status.png)](https://godoc.org/github.com/lancer-kit/armory)
[![Go Report Card](https://goreportcard.com/badge/github.com/lancer-kit/armory)](https://goreportcard.com/report/github.com/lancer-kit/armory)

# Armory Service Kit

Common libraries for building go services:

- **Api**
    - [Render](./api/render/README.md) - response helper, base responses
    - [HTTPX](./api/httpx) - wrapper for `http.Client` with additional helpers the for RESTfull APIs. 
- [DB](./db/README.md) - connector for the ORMless interaction with the PostgreSQL databases.
- [Log](./log/README.md) - simple wrapper for logrus with some useful perks.

- [Auth](./auth/README.md) - methods for the service authorization.
- [Crypto](./crypto/README.md) - wrappers for hashing, signing, random values generation etc.
- [natsx](./natsx/README.md) - simple wrapper for NATS.

