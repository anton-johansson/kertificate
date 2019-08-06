# Kertificate

Kertificate is a PKI management system that allows you to administrate your common authorities and server and client certificates.


## Development

Start API server (runs on port `8080`):

```shell
$ make linux && ./bin/kertificate-linux-amd64 start
```

Start web UI with hot-reload (runs on port `8000`):

```shell
$ cd web/
$ npm start
```


## Adding license notice to files

```shell
$ go get -u github.com/google/addlicense
$ addlicense -c 'Anton Johansson' *
```

Note: Above command might only work for ZSH. Need to verify and fix so we can verify this in CI.


## License

This tool is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
