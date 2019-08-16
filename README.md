![Kert](./resources/kert.svg)

<h1 align="center">Kertificate</h1>
<p align="center">
    <i>Kert loves public key infrastructure!</i>
</p>
<p align="center">
    Kertificate is a PKI management system that allows you to administrate your common authorities and server and client certificates and perform actions upon renewal of these certificates.
</p>
<p align="center">
    <a aria-label="build status" href="https://circleci.com/gh/anton-johansson/workflows/kertificate">
        <img src="https://img.shields.io/circleci/build/gh/anton-johansson/mattermost-housekeeper?style=for-the-badge">
    </a>
    <a aria-label="contributors graph" href="https://github.com/anton-johansson/kertificate/graphs/contributors">
        <img src="https://img.shields.io/github/contributors/anton-johansson/kertificate.svg?style=for-the-badge">
    </a>
    <a aria-label="license" href="https://github.com/anton-johansson/kertificate/blob/add-logo/LICENSE">
        <img src="https://img.shields.io/github/license/anton-johansson/kertificate.svg?style=for-the-badge">
    </a>
</p>


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
