FROM golang:latest AS build

WORKDIR /opt
COPY main.go /opt/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch
COPY --from=0 /opt/main /main
COPY filler.jpg /filler.jpg

ENV PORT 80

EXPOSE 80/tcp

ENTRYPOINT ["/main"]
