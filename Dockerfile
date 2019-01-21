FROM golang:1.11.4 as base
WORKDIR /tmp/gateway
COPY . .
RUN go build -mod=vendor -o /tmp/service .

FROM alpine:latest
WORKDIR /tmp
COPY --from=base /tmp/service .
ENTRYPOINT ./service --service_port=80
EXPOSE 80
