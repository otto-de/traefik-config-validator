FROM golang:1.17 as builder

WORKDIR /opt/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go test -cover ./...

ENV GOOS linux
ARG GOARCH
ENV GOARCH ${GOARCH:-amd64}
ENV CGO_ENABLED=0

ARG VERSION
RUN go build -v -ldflags "-X main.version=$VERSION" -o traefik-config-validator cmd/traefik-config-validator/main.go

FROM alpine:latest

ENTRYPOINT ["traefik-config-validator"]
CMD ["--config.file=/tmp/config.yml"]
RUN addgroup -g 1000 validator && \
    adduser -u 1000 -D -G validator validator -h /validator

WORKDIR /validator/

RUN apk --no-cache add ca-certificates
COPY --from=builder /opt/traefik-config-validator /usr/local/bin/traefik-config-validator
USER validator
