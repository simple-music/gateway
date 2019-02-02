FROM golang:1.11.4 as base
WORKDIR /tmp/gateway
COPY . .
RUN go build -mod=vendor -o /tmp/service .

FROM ubuntu:latest
WORKDIR /tmp
COPY --from=base /tmp/service ./service
CMD ./service --service_port=80
EXPOSE 80
