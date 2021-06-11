FROM golang:latest

RUN mkdir /go_videos

COPY . /go_videos/

WORKDIR /go_videos

RUN go get -d -v ./...

RUN go build

# Exposing the default port
EXPOSE 8000

CMD go run /go_videos/main.go