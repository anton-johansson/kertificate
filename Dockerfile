# Copyright 2019 Anton Johansson
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM node:12.8.0-slim as build-web
RUN apt update && apt install make
WORKDIR /src/app
COPY . /src/app
RUN make build-web

FROM golang:1.12.2-alpine3.9 AS build-go
RUN apk add --no-cache git make
WORKDIR /src/app
COPY . /src/app
COPY --from=build-web /src/app/web/dist /src/app/web/dist
RUN make linux

FROM alpine:3.9.3
COPY --from=build-go /src/app/bin/kertificate-linux-amd64 /usr/bin/kertificate
RUN addgroup -S -g 1000 kertificate && adduser -S -u 1000 -G kertificate kertificate
USER 1000
ENTRYPOINT ["kertificate", "start"]
