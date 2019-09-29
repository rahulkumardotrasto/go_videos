FROM golang:latest

RUN mkdir /gt_node_backend

COPY . /gt_node_backend/

WORKDIR /gt_node_backend

RUN go get -d -v ./...

RUN go build

# Exposing the default port
EXPOSE 8000

CMD go run /gt_node_backend/main.go