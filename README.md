# PKI management system

PKIMS is a web UI that allows you to administrate your common authorities and server certificates.


## Development

Start API server (runs on port `8080`):

```shell
$ make linux && ./bin/pkims-linux-amd64 start
```

Start web UI with hot-reload (runs on port `8000`):

```shell
$ cd web/
$ npm start
```
