FROM shvid/ubuntu-golang as builder

ARG VERSION
ARG BUILD

WORKDIR /go/src/github.com/arpabet/sprint
ADD . .

RUN go build -o /sprint -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

FROM ubuntu:18.04
WORKDIR /app

COPY --from=builder /sprint .

EXPOSE 7000 7080

CMD ["/app/sprint", "start"]

