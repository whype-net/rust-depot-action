FROM golang:1.23-alpine as builder

WORKDIR /build
COPY . .
ENV CGO_ENABLE=0
RUN go build -o /build/builder .

FROM alpine:3.20

WORKDIR /srv
COPY --from=builder /build/builder .
ENTRYPOINT ["/srv/builder"]