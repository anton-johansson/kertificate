# Kertificate

Kertificate is a PKI management system that allows you to administrate your common authorities and server and client certificates.


## Development

Run the API server - with files watched - on port `8080` and the web application with hot reload on port `8000`. This requires `make`, `go` version 1.12 or higher, `node` and `npm` installed.

```shell
$ make -j 2
```


## Installing

To create a production-ready build (requires `make`, `go` version 1.12 or higher, `node` and `npm` installed):

```
$ make build-linux
```


### Docker

To create a production-ready Docker image (which only requires `make` and `docker` installed):

```
$ make docker
```


## Adding license notice to files

```shell
$ go get -u github.com/google/addlicense
$ addlicense -c 'Anton Johansson' *
```

Note: Above command might only work for ZSH. Need to verify and fix so we can verify this in CI.


## License

This tool is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
