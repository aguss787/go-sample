FROM golang:latest AS builder

WORKDIR /opt

COPY . .
RUN ["go", "build"]

FROM ubuntu:23.10 AS runner

RUN apt-get update
RUN apt-get install -y libpq5

WORKDIR /srv
COPY var/development.yaml var/development.yaml
COPY --from=builder /opt/api glints-api

CMD ["./glints-api"]
