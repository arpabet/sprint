# Copyright (c) Karagatan LLC.
# SPDX-License-Identifier: Apache-2.0

FROM codeallergy/ubuntu-golang as builder

ARG VERSION
ARG BUILD

WORKDIR /go/src/go.arpabet.com/sprint/framework
ADD . .

ENV GONOSUMDB github.com

RUN go build -o /sprint -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

FROM ubuntu:18.04
WORKDIR /app/bin

COPY --from=builder /sprint .

EXPOSE 8080 8443 8444

CMD ["/app/bin/sprint", "run"]

