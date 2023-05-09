FROM golang:1.18-alpine as builder
WORKDIR /build
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY . .
COPY src .
RUN go build -o /main src/cmd/main.go

FROM alpine:latest
COPY --from=builder main /bin/main
COPY src/config/config.json /config/config.json
ENTRYPOINT ["/bin/main"]