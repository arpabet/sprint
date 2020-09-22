FROM shvid/ubuntu-golang as builder

ARG VERSION
ARG BUILD

WORKDIR /go/src/github.com/arpabet/templateserv
ADD . .

RUN sed -i "s/%TAG%/${TAG}/g" main.go && \
    go build -o /templateserv -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

FROM ubuntu:18.04
WORKDIR /app

COPY --from=builder /templateserv .

EXPOSE 8433 8434

CMD ["/app/templateserv", "start"]

