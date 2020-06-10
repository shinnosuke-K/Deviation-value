FROM golang:1.13.4 as build
WORKDIR /go/src
COPY $PWD/ /go/src
RUN CGO_ENABLED=0 go build -o server server.go
CMD ["go", "run", "server.go"]

#FROM ubuntu:latest
#COPY --from=build /go/src /go/src
#WORKDIR /go/src
#CMD ["go", "run", "server.go"]