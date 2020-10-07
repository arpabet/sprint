FROM shvid/ubuntu-golang as builder

ARG VERSION
ARG BUILD

WORKDIR /go/src/github.com/arpabet/sprint
ADD . .

RUN sed -i "s/%TAG%/${TAG}/g" main.go && \
    go build -o /sprint -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

FROM ubuntu:18.04
WORKDIR /app

COPY --from=builder /sprint .

EXPOSE 8433 8434

CMD ["/app/sprint", "start"]

