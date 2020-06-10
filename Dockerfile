FROM golang:1.13.4 as build
WORKDIR /go/src
COPY $PWD/ /go/src
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server server.go

FROM alpine:latest
COPY --from=build /go/src /go/src
WORKDIR /go/src
CMD ["/go/src/server"]