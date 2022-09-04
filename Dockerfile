FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update

# build go app
RUN go mod download
RUN go build -o notes-log ./cmd/main.go

CMD ["./notes-log"]