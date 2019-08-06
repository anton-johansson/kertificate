FROM golang:1.12.2-alpine3.9 AS build
RUN apk add --no-cache git make
COPY . /src/golang
WORKDIR /src/golang
RUN make linux

FROM alpine:3.9.3
COPY --from=build /src/golang/build/kertificate-linux-amd64 /usr/bin/kertificate
RUN addgroup -S -g 1000 kertificate && adduser -S -u 1000 -G kertificate kertificate
USER 1000
ENTRYPOINT ["kertificate", "start"]
