FROM golang:1.24.5-bookworm AS builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
WORKDIR /app
COPY ./ ./
RUN go mod download && go build -o /app/out/incident-notification /app/main.go


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder --chown=nobody:nogroup /app/out/incident-notification /app/

ENTRYPOINT ["/app/incident-notification"]
CMD ["--help"]
