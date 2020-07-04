FROM shvid/ubuntu-golang as builder

ARG TAG

WORKDIR /go/src/github.com/arpabet/template-server
ADD . .

RUN sed -i "s/%TAG%/${TAG}/g" main.go && \
    go build -o /template-server

FROM ubuntu:18.04
WORKDIR /app

COPY --from=builder /template-server .

EXPOSE 8080 8081

CMD ["/app/template-server"]

